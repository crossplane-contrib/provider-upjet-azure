// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	hadoopcluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/hadoopcluster"
	hbasecluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/hbasecluster"
	interactivequerycluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/interactivequerycluster"
	kafkacluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/kafkacluster"
	sparkcluster "github.com/upbound/provider-azure/internal/controller/cluster/hdinsight/sparkcluster"
)

// Setup_hdinsight creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_hdinsight(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		hadoopcluster.Setup,
		hbasecluster.Setup,
		interactivequerycluster.Setup,
		kafkacluster.Setup,
		sparkcluster.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
