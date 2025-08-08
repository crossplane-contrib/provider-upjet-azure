// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	api "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/api"
	apidiagnostic "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apidiagnostic"
	apioperation "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apioperation"
	apioperationpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apioperationpolicy"
	apioperationtag "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apioperationtag"
	apipolicy "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apipolicy"
	apirelease "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apirelease"
	apischema "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apischema"
	apitag "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apitag"
	apiversionset "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/apiversionset"
	authorizationserver "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/authorizationserver"
	backend "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/backend"
	certificate "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/certificate"
	customdomain "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/customdomain"
	diagnostic "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/diagnostic"
	emailtemplate "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/emailtemplate"
	gateway "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/gateway"
	gatewayapi "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/gatewayapi"
	globalschema "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/globalschema"
	identityprovideraad "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/identityprovideraad"
	identityproviderfacebook "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/identityproviderfacebook"
	identityprovidergoogle "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/identityprovidergoogle"
	identityprovidermicrosoft "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/identityprovidermicrosoft"
	identityprovidertwitter "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/identityprovidertwitter"
	logger "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/logger"
	management "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/management"
	namedvalue "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/namedvalue"
	notificationrecipientemail "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/notificationrecipientemail"
	notificationrecipientuser "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/notificationrecipientuser"
	openidconnectprovider "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/openidconnectprovider"
	policy "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/policy"
	product "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/product"
	productapi "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/productapi"
	productpolicy "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/productpolicy"
	producttag "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/producttag"
	rediscache "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/rediscache"
	subscription "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/subscription"
	tag "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/tag"
	user "github.com/upbound/provider-azure/internal/controller/namespaced/apimanagement/user"
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
		customdomain.Setup,
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

// SetupGated_apimanagement creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_apimanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		api.SetupGated,
		apidiagnostic.SetupGated,
		apioperation.SetupGated,
		apioperationpolicy.SetupGated,
		apioperationtag.SetupGated,
		apipolicy.SetupGated,
		apirelease.SetupGated,
		apischema.SetupGated,
		apitag.SetupGated,
		apiversionset.SetupGated,
		authorizationserver.SetupGated,
		backend.SetupGated,
		certificate.SetupGated,
		customdomain.SetupGated,
		diagnostic.SetupGated,
		emailtemplate.SetupGated,
		gateway.SetupGated,
		gatewayapi.SetupGated,
		globalschema.SetupGated,
		identityprovideraad.SetupGated,
		identityproviderfacebook.SetupGated,
		identityprovidergoogle.SetupGated,
		identityprovidermicrosoft.SetupGated,
		identityprovidertwitter.SetupGated,
		logger.SetupGated,
		management.SetupGated,
		namedvalue.SetupGated,
		notificationrecipientemail.SetupGated,
		notificationrecipientuser.SetupGated,
		openidconnectprovider.SetupGated,
		policy.SetupGated,
		product.SetupGated,
		productapi.SetupGated,
		productpolicy.SetupGated,
		producttag.SetupGated,
		rediscache.SetupGated,
		subscription.SetupGated,
		tag.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
