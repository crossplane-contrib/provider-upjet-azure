// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	springcloudaccelerator "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudaccelerator"
	springcloudactivedeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudactivedeployment"
	springcloudapiportal "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudapiportal"
	springcloudapiportalcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudapiportalcustomdomain"
	springcloudapp "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudapp"
	springcloudappcosmosdbassociation "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudappcosmosdbassociation"
	springcloudappmysqlassociation "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudappmysqlassociation"
	springcloudappredisassociation "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudappredisassociation"
	springcloudbuilddeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudbuilddeployment"
	springcloudbuilder "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudbuilder"
	springcloudbuildpackbinding "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudbuildpackbinding"
	springcloudcertificate "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcertificate"
	springcloudconfigurationservice "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudconfigurationservice"
	springcloudcontainerdeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcontainerdeployment"
	springcloudcustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcustomdomain"
	springcloudcustomizedaccelerator "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudcustomizedaccelerator"
	springclouddevtoolportal "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springclouddevtoolportal"
	springcloudgateway "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudgateway"
	springcloudgatewaycustomdomain "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudgatewaycustomdomain"
	springcloudjavadeployment "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudjavadeployment"
	springcloudservice "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudservice"
	springcloudstorage "github.com/upbound/provider-azure/internal/controller/cluster/appplatform/springcloudstorage"
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
