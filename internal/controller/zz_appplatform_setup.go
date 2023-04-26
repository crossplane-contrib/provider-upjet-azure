/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	springcloudaccelerator "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudaccelerator"
	springcloudactivedeployment "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudactivedeployment"
	springcloudapiportal "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudapiportal"
	springcloudapiportalcustomdomain "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudapiportalcustomdomain"
	springcloudapp "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudapp"
	springcloudappcosmosdbassociation "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudappredisassociation"
	springcloudbuilddeployment "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudbuilddeployment"
	springcloudbuilder "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudbuilder"
	springcloudbuildpackbinding "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudbuildpackbinding"
	springcloudcertificate "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudcertificate"
	springcloudconfigurationservice "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudconfigurationservice"
	springcloudcontainerdeployment "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudcontainerdeployment"
	springcloudcustomdomain "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudcustomdomain"
	springcloudcustomizedaccelerator "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudcustomizedaccelerator"
	springclouddevtoolportal "github.com/upbound/provider-azure/internal/controller/appplatform/springclouddevtoolportal"
	springcloudgateway "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudgateway"
	springcloudgatewaycustomdomain "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudgatewaycustomdomain"
	springcloudjavadeployment "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudjavadeployment"
	springcloudservice "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudservice"
	springcloudstorage "github.com/upbound/provider-azure/internal/controller/appplatform/springcloudstorage"
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
