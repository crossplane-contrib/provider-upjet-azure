// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	hadoopcluster "github.com/upbound/provider-azure/v2/internal/controller/namespaced/hdinsight/hadoopcluster"
	hbasecluster "github.com/upbound/provider-azure/v2/internal/controller/namespaced/hdinsight/hbasecluster"
	interactivequerycluster "github.com/upbound/provider-azure/v2/internal/controller/namespaced/hdinsight/interactivequerycluster"
	kafkacluster "github.com/upbound/provider-azure/v2/internal/controller/namespaced/hdinsight/kafkacluster"
	sparkcluster "github.com/upbound/provider-azure/v2/internal/controller/namespaced/hdinsight/sparkcluster"
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

// SetupGated_hdinsight creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_hdinsight(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		hadoopcluster.SetupGated,
		hbasecluster.SetupGated,
		interactivequerycluster.SetupGated,
		kafkacluster.SetupGated,
		sparkcluster.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
