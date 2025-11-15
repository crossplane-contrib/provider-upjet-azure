// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	springcloudaccelerator "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudaccelerator"
	springcloudactivedeployment "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudactivedeployment"
	springcloudapiportal "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudapiportal"
	springcloudapiportalcustomdomain "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudapiportalcustomdomain"
	springcloudapp "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudapp"
	springcloudappcosmosdbassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudappredisassociation"
	springcloudbuilddeployment "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudbuilddeployment"
	springcloudbuilder "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudbuilder"
	springcloudbuildpackbinding "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudbuildpackbinding"
	springcloudcertificate "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudcertificate"
	springcloudconfigurationservice "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudconfigurationservice"
	springcloudcontainerdeployment "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudcontainerdeployment"
	springcloudcustomdomain "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudcustomdomain"
	springcloudcustomizedaccelerator "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudcustomizedaccelerator"
	springclouddevtoolportal "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springclouddevtoolportal"
	springcloudgateway "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudgateway"
	springcloudgatewaycustomdomain "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudgatewaycustomdomain"
	springcloudjavadeployment "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudjavadeployment"
	springcloudservice "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudservice"
	springcloudstorage "github.com/upbound/provider-azure/v2/internal/controller/namespaced/appplatform/springcloudstorage"
)

// Setup_appplatform creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		springcloudaccelerator.Setup,
		springcloudactivedeployment.Setup,
		springcloudapiportal.Setup,
		springcloudapiportalcustomdomain.Setup,
		springcloudapp.Setup,
		springcloudappcosmosdbassociation.Setup,
		springcloudappmysqlassociation.Setup,
		springcloudappredisassociation.Setup,
		springcloudbuilddeployment.Setup,
		springcloudbuilder.Setup,
		springcloudbuildpackbinding.Setup,
		springcloudcertificate.Setup,
		springcloudconfigurationservice.Setup,
		springcloudcontainerdeployment.Setup,
		springcloudcustomdomain.Setup,
		springcloudcustomizedaccelerator.Setup,
		springclouddevtoolportal.Setup,
		springcloudgateway.Setup,
		springcloudgatewaycustomdomain.Setup,
		springcloudjavadeployment.Setup,
		springcloudservice.Setup,
		springcloudstorage.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_appplatform creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_appplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		springcloudaccelerator.SetupGated,
		springcloudactivedeployment.SetupGated,
		springcloudapiportal.SetupGated,
		springcloudapiportalcustomdomain.SetupGated,
		springcloudapp.SetupGated,
		springcloudappcosmosdbassociation.SetupGated,
		springcloudappmysqlassociation.SetupGated,
		springcloudappredisassociation.SetupGated,
		springcloudbuilddeployment.SetupGated,
		springcloudbuilder.SetupGated,
		springcloudbuildpackbinding.SetupGated,
		springcloudcertificate.SetupGated,
		springcloudconfigurationservice.SetupGated,
		springcloudcontainerdeployment.SetupGated,
		springcloudcustomdomain.SetupGated,
		springcloudcustomizedaccelerator.SetupGated,
		springclouddevtoolportal.SetupGated,
		springcloudgateway.SetupGated,
		springcloudgatewaycustomdomain.SetupGated,
		springcloudjavadeployment.SetupGated,
		springcloudservice.SetupGated,
		springcloudstorage.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
