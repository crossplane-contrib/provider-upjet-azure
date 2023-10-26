// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	api "github.com/upbound/provider-azure/internal/controller/apimanagement/api"
	apidiagnostic "github.com/upbound/provider-azure/internal/controller/apimanagement/apidiagnostic"
	apioperation "github.com/upbound/provider-azure/internal/controller/apimanagement/apioperation"
	apioperationpolicy "github.com/upbound/provider-azure/internal/controller/apimanagement/apioperationpolicy"
	apioperationtag "github.com/upbound/provider-azure/internal/controller/apimanagement/apioperationtag"
	apipolicy "github.com/upbound/provider-azure/internal/controller/apimanagement/apipolicy"
	apirelease "github.com/upbound/provider-azure/internal/controller/apimanagement/apirelease"
	apischema "github.com/upbound/provider-azure/internal/controller/apimanagement/apischema"
	apitag "github.com/upbound/provider-azure/internal/controller/apimanagement/apitag"
	apiversionset "github.com/upbound/provider-azure/internal/controller/apimanagement/apiversionset"
	authorizationserver "github.com/upbound/provider-azure/internal/controller/apimanagement/authorizationserver"
	backend "github.com/upbound/provider-azure/internal/controller/apimanagement/backend"
	certificate "github.com/upbound/provider-azure/internal/controller/apimanagement/certificate"
	diagnostic "github.com/upbound/provider-azure/internal/controller/apimanagement/diagnostic"
	emailtemplate "github.com/upbound/provider-azure/internal/controller/apimanagement/emailtemplate"
	gateway "github.com/upbound/provider-azure/internal/controller/apimanagement/gateway"
	gatewayapi "github.com/upbound/provider-azure/internal/controller/apimanagement/gatewayapi"
	globalschema "github.com/upbound/provider-azure/internal/controller/apimanagement/globalschema"
	identityprovideraad "github.com/upbound/provider-azure/internal/controller/apimanagement/identityprovideraad"
	identityproviderfacebook "github.com/upbound/provider-azure/internal/controller/apimanagement/identityproviderfacebook"
	identityprovidergoogle "github.com/upbound/provider-azure/internal/controller/apimanagement/identityprovidergoogle"
	identityprovidermicrosoft "github.com/upbound/provider-azure/internal/controller/apimanagement/identityprovidermicrosoft"
	identityprovidertwitter "github.com/upbound/provider-azure/internal/controller/apimanagement/identityprovidertwitter"
	logger "github.com/upbound/provider-azure/internal/controller/apimanagement/logger"
	management "github.com/upbound/provider-azure/internal/controller/apimanagement/management"
	namedvalue "github.com/upbound/provider-azure/internal/controller/apimanagement/namedvalue"
	notificationrecipientemail "github.com/upbound/provider-azure/internal/controller/apimanagement/notificationrecipientemail"
	notificationrecipientuser "github.com/upbound/provider-azure/internal/controller/apimanagement/notificationrecipientuser"
	openidconnectprovider "github.com/upbound/provider-azure/internal/controller/apimanagement/openidconnectprovider"
	policy "github.com/upbound/provider-azure/internal/controller/apimanagement/policy"
	product "github.com/upbound/provider-azure/internal/controller/apimanagement/product"
	productapi "github.com/upbound/provider-azure/internal/controller/apimanagement/productapi"
	productpolicy "github.com/upbound/provider-azure/internal/controller/apimanagement/productpolicy"
	producttag "github.com/upbound/provider-azure/internal/controller/apimanagement/producttag"
	rediscache "github.com/upbound/provider-azure/internal/controller/apimanagement/rediscache"
	subscription "github.com/upbound/provider-azure/internal/controller/apimanagement/subscription"
	tag "github.com/upbound/provider-azure/internal/controller/apimanagement/tag"
	user "github.com/upbound/provider-azure/internal/controller/apimanagement/user"
)

// Setup_apimanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apimanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		api.Setup,
		apidiagnostic.Setup,
		apioperation.Setup,
		apioperationpolicy.Setup,
		apioperationtag.Setup,
		apipolicy.Setup,
		apirelease.Setup,
		apischema.Setup,
		apitag.Setup,
		apiversionset.Setup,
		authorizationserver.Setup,
		backend.Setup,
		certificate.Setup,
		diagnostic.Setup,
		emailtemplate.Setup,
		gateway.Setup,
		gatewayapi.Setup,
		globalschema.Setup,
		identityprovideraad.Setup,
		identityproviderfacebook.Setup,
		identityprovidergoogle.Setup,
		identityprovidermicrosoft.Setup,
		identityprovidertwitter.Setup,
		logger.Setup,
		management.Setup,
		namedvalue.Setup,
		notificationrecipientemail.Setup,
		notificationrecipientuser.Setup,
		openidconnectprovider.Setup,
		policy.Setup,
		product.Setup,
		productapi.Setup,
		productpolicy.Setup,
		producttag.Setup,
		rediscache.Setup,
		subscription.Setup,
		tag.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
