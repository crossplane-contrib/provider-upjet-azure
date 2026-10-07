// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfazureclient "github.com/hashicorp/terraform-provider-azurerm/xpprovider"
)

// armHost is the Azure Resource Manager endpoint of the public cloud, which is
// the environment an offline configuration with no ProviderConfig defaults to.
const armHost = "management.azure.com"

// graphHost is the Microsoft Graph endpoint. clients.Build authorizes against
// it (not armHost) purely to inspect the token's claims; see
// clients.NewResourceManagerAccount.
const graphHost = "graph.microsoft.com"

// loginHost is the endpoint the Azure SDK requests access tokens from.
const loginHost = "login.microsoftonline.com"

// realAuthClient is the Azure SDK's own token HTTP client, captured during
// package variable initialisation - that is, before any test can call
// EnableOfflineAuthentication, whose sync.Once swap cannot be undone. Holding
// on to it lets TestConfigureEgressesWithoutOfflineAuthentication put the real
// client back regardless of what ran first.
var realAuthClient = auth.Client

// egress is the tripwire armed by TestMain, shared by every test in the
// package. Tests that read it must not run in parallel with one another.
var egress *proxyRecorder

// proxyRecorder is an HTTP proxy that records every request routed through it
// and forwards none of them.
//
// Every HTTP transport the Azure SDK builds - the one that fetches access
// tokens (go-azure-sdk/sdk/auth), the one that talks to ARM and Microsoft
// Graph (go-azure-sdk/sdk/client), and the ad hoc one clients/graph builds for
// claims-based oid discovery - construct their own http.Transport with
// Proxy: http.ProxyFromEnvironment rather than using http.DefaultTransport.
// Pointing HTTPS_PROXY at a recorder therefore puts a tripwire on the
// process's real network boundary: nothing is stubbed or injected, and a
// request cannot reach Azure without being recorded here first. The same
// tripwire works unchanged against the compiled provider binary.
type proxyRecorder struct {
	url string

	mu   sync.Mutex
	seen []string
}

func (p *proxyRecorder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CONNECT carries its target in Host; an absolute-form proxy request
	// carries it in URL.
	target := r.Host
	if target == "" && r.URL != nil {
		target = r.URL.Host
	}
	p.mu.Lock()
	p.seen = append(p.seen, r.Method+" "+target)
	p.mu.Unlock()
	http.Error(w, "blocked by the egress tripwire", http.StatusBadGateway)
}

func (p *proxyRecorder) records() []string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return append([]string(nil), p.seen...)
}

func (p *proxyRecorder) reset() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.seen = nil
}

// reached reports whether any recorded request targeted host.
func (p *proxyRecorder) reached(host string) bool {
	for _, r := range p.records() {
		if strings.Contains(r, host) {
			return true
		}
	}
	return false
}

func TestMain(m *testing.M) {
	os.Exit(runTests(m))
}

func runTests(m *testing.M) int {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot listen for the egress tripwire: %v\n", err)
		return 1
	}
	defer l.Close() //nolint:errcheck // the process is exiting

	egress = &proxyRecorder{url: "http://" + l.Addr().String()}
	srv := &http.Server{Handler: egress, ReadHeaderTimeout: 5 * time.Second}
	go srv.Serve(l)   //nolint:errcheck // Serve always returns a non-nil error
	defer srv.Close() //nolint:errcheck // the process is exiting

	// net/http resolves the proxy environment exactly once per process, so
	// this has to happen before anything in the package makes a request.
	os.Setenv("HTTP_PROXY", egress.url)  //nolint:errcheck // must outlive any single test, so t.Setenv will not do
	os.Setenv("HTTPS_PROXY", egress.url) //nolint:errcheck // must outlive any single test, so t.Setenv will not do
	os.Setenv("NO_PROXY", "")            //nolint:errcheck // must outlive any single test, so t.Setenv will not do

	return m.Run()
}

// Ordering note: two tests below (WithoutObjectIDClaimEgresses and
// UnguardedARMCall) deliberately run a real, unguarded Azure call that the
// tripwire answers with a 502 - a status the SDK's retry policy treats as
// retryable - against a context neither test can give a deadline (explained
// at each one). Rather than block on that retry loop, each starts it in a
// background goroutine and only waits for the tripwire's first recording,
// leaving the goroutine to keep retrying in the background until the test
// binary exits. Every test that asserts zero egress is placed *before* both
// of these in this file, so no leaked retry from one of them can land a stray
// record in a later test's window. Keep that order if you add more tests.

// TestOfflineConfigureDoesNotEgress asserts that configuring the AzureRM
// Terraform provider for offline diffs makes no outbound request whatsoever.
//
// This is worth asserting because the egress guard is installed on the
// provider's Meta only *after* p.Configure returns - see configureOffline -
// so anything Configure itself calls goes out unguarded. clients.Build does
// contain two such calls, detailed in configureOffline's own comment; both are
// closed by configuration or by the offline token's claims, not by the guard.
func TestOfflineConfigureDoesNotEgress(t *testing.T) {
	EnableOfflineAuthentication()
	egress.reset()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureRM provider schema: %v", err)
	}
	if _, err := configureOffline(context.Background(), p, nil); err != nil {
		t.Fatalf("cannot configure the AzureRM provider offline: %v", err)
	}

	if got := egress.records(); len(got) > 0 {
		t.Errorf("configuring the provider offline reached the network: %v", got)
	}
}

// TestConfigureEgressesWithoutOfflineAuthentication is the positive control for
// TestOfflineConfigureDoesNotEgress at the Configure boundary specifically.
// Putting the Azure SDK's real token client back makes Configure attempt a
// genuine token request, which the tripwire records against loginHost. That is
// what makes the silence in TestOfflineConfigureDoesNotEgress meaningful: the
// tripwire does watch this code path, and EnableOfflineAuthentication is what
// keeps it quiet.
//
// This one blocks rather than backgrounding: the token client's own retry
// count is bounded regardless of context deadline, so it returns in seconds,
// not minutes.
func TestConfigureEgressesWithoutOfflineAuthentication(t *testing.T) {
	EnableOfflineAuthentication()
	t.Cleanup(func() { auth.Client = offlineTokenClient{} })
	auth.Client = realAuthClient
	egress.reset()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureRM provider schema: %v", err)
	}
	if _, err := configureOffline(context.Background(), p, nil); err == nil {
		t.Fatal("expected configuring the provider to fail without the offline token client")
	}

	if !egress.reached(loginHost) {
		t.Errorf("expected Configure to request a token from %s, recorded: %v", loginHost, egress.records())
	}
}

// refreshResourceGroup runs the real azurerm_resource_group Read against ps's
// configured client - the same call a diff server would make to populate the
// actual resource's Terraform state when status.atProvider is observed but
// incomplete. It is exactly what a test wants here: a real API call, through a
// real, Build()-registered ARM client, with no CustomizeDiff of its own to
// complicate the picture.
func refreshResourceGroup(ctx context.Context, p *schema.Provider, meta any) error {
	r := p.ResourcesMap[azurermResourceGroup]
	state := &tfsdk.InstanceState{
		ID: "/subscriptions/" + offlineSubscriptionID + "/resourceGroups/demo-rg",
		Attributes: map[string]string{
			"id":   "/subscriptions/" + offlineSubscriptionID + "/resourceGroups/demo-rg",
			"name": "demo-rg",
		},
	}
	_, diags := r.RefreshWithoutUpgrade(ctx, state, meta)
	if diags.HasError() {
		return fmt.Errorf("%v", diags)
	}
	return nil
}

// TestOfflineReadBlockedByGuard asserts that a real Read - the same call
// TestEgressTripwireObservesUnguardedARMCall below proves reaches Azure when
// unguarded - is refused immediately rather than reaching the network when
// run through the full configureOffline path, guard installed.
func TestOfflineReadBlockedByGuard(t *testing.T) {
	EnableOfflineAuthentication()
	egress.reset()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureRM provider schema: %v", err)
	}
	ps, err := configureOffline(context.Background(), p, nil)
	if err != nil {
		t.Fatalf("cannot configure the AzureRM provider offline: %v", err)
	}

	err = refreshResourceGroup(context.Background(), p, ps.Meta)
	if err == nil {
		t.Fatal("expected the offline Read to fail, the guard must refuse the ARM call")
	}
	if !strings.Contains(err.Error(), "offline diff server blocked an outbound Azure API request") {
		t.Errorf("expected the request guard's error, got: %v", err)
	}
	if got := egress.records(); len(got) > 0 {
		t.Errorf("the guarded Read reached the network: %v", got)
	}
}

// noObjectIDTokenClient answers token requests with an otherwise valid offline
// token that carries no oid claim.
type noObjectIDTokenClient struct{}

func (noObjectIDTokenClient) Do(req *http.Request) (*http.Response, error) {
	payload, err := json.Marshal(map[string]any{
		"aud":   "https://" + graphHost,
		"iss":   "https://sts.windows.net/" + offlineTenantID + "/",
		"tid":   offlineTenantID,
		"appid": offlineClientID,
		// An empty scp makes NewResourceManagerAccount treat this as a
		// service principal and take the ServicePrincipalObjectID branch;
		// both branches reach Microsoft Graph.
		"scp": "",
	})
	if err != nil {
		return nil, err
	}
	enc := base64.RawURLEncoding.EncodeToString
	body, err := json.Marshal(map[string]any{
		"access_token": enc([]byte(`{"alg":"none","typ":"JWT"}`)) + "." + enc(payload) + ".offline",
		"token_type":   "Bearer",
		"expires_in":   offlineTokenLifetime,
	})
	if err != nil {
		return nil, err
	}

	return &http.Response{
		Status:        "200 OK",
		StatusCode:    http.StatusOK,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        http.Header{"Content-Type": []string{"application/json"}},
		Body:          io.NopCloser(strings.NewReader(string(body))),
		ContentLength: int64(len(body)),
		Request:       req,
	}, nil
}

// TestOfflineConfigureWithoutObjectIDClaimEgresses strips the oid claim from
// the offline access token to reach clients.NewResourceManagerAccount's
// Microsoft Graph fallback lookup, and shows that - unlike the equivalent gap
// in a sibling provider sharing this same Azure SDK, where a missing context
// deadline independently blocked the call even once reached - this one really
// does leave the process.
//
// The reason is in clients/graph: ServicePrincipalObjectID and
// UserPrincipalObjectID each synthesize their own 5-minute deadline when the
// incoming context carries none, rather than refusing to proceed. They also
// build their own ad hoc *msgraph.Client via graphClient(), never adding it to
// the client.armClients slice AppendRequestMiddleware iterates - so even with
// configureOffline's guard installed, it would not have covered this client
// regardless of timing.
//
// What actually keeps this path closed, today, is a single fact pinned
// elsewhere: offlineAccessToken always mints a non-empty oid. This test exists
// so that if that ever stopped being true, something here would fail loudly
// instead of the gap being reintroduced silently.
//
// It runs the real call in the background rather than waiting for it: see the
// ordering note above this file's first test.
func TestOfflineConfigureWithoutObjectIDClaimEgresses(t *testing.T) {
	EnableOfflineAuthentication()
	t.Cleanup(func() { auth.Client = offlineTokenClient{} })
	auth.Client = noObjectIDTokenClient{}
	egress.reset()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureRM provider schema: %v", err)
	}
	go func() {
		//nolint:errcheck // the call is expected to fail; what matters is that it was seen leaving
		configureOffline(context.Background(), p, nil)
	}()

	deadline := time.Now().Add(10 * time.Second)
	for !egress.reached(graphHost) && time.Now().Before(deadline) {
		time.Sleep(50 * time.Millisecond)
	}
	if !egress.reached(graphHost) {
		t.Errorf("expected the oid-discovery fallback to reach %s without the oid claim, recorded: %v", graphHost, egress.records())
	}
}

// TestEgressTripwireObservesUnguardedARMCall is the positive control for the
// guard mechanism itself, independent of any one resource's CustomizeDiff: it
// configures the provider exactly as configureOffline does but *without*
// installing denyOutboundRequest, then runs a real Read against a
// Build()-registered ARM client. With nothing in the way, the request reaches
// the network boundary and the tripwire records it.
//
// resourceResourceGroupRead (the legacy, no-context Read signature) derives
// its own request deadline from meta.(*clients.Client).StopContext - fixed at
// Configure time via timeouts.ForRead, and already stripped of any deadline by
// configureNoForkAzureClient's context.WithoutCancel - not from whatever
// context RefreshWithoutUpgrade is called with. So nothing this test passes in
// can bound how long the real call keeps retrying against the tripwire's
// 502s: it runs out to the resource's own 5-minute default Read timeout
// regardless. It runs the real call in the background rather than waiting for
// it for the same reason as the test above; see the ordering note at the top
// of this file. This call path is purely a test device for proving the guard
// works - the real diff server never takes it, since it only ever calls
// CustomizeDiff, whose context genuinely is threaded through.
func TestEgressTripwireObservesUnguardedARMCall(t *testing.T) {
	EnableOfflineAuthentication()
	egress.reset()

	p, err := tfazureclient.GetProviderSchema(context.Background())
	if err != nil {
		t.Fatalf("cannot get the AzureRM provider schema: %v", err)
	}
	ps := &terraform.Setup{Configuration: offlineConfiguration(nil)}
	if err := configureNoForkAzureClient(context.Background(), ps, *p); err != nil {
		t.Fatalf("cannot configure the AzureRM provider: %v", err)
	}

	go func() {
		//nolint:errcheck // the call is expected to fail; what matters is that it was seen leaving
		refreshResourceGroup(context.Background(), p, ps.Meta)
	}()

	deadline := time.Now().Add(10 * time.Second)
	for !egress.reached(armHost) && time.Now().Before(deadline) {
		time.Sleep(50 * time.Millisecond)
	}
	if !egress.reached(armHost) {
		t.Errorf("expected the unguarded Read to reach %s, recorded: %v", armHost, egress.records())
	}
}
