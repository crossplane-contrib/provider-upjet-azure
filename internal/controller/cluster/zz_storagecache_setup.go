// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	hpccache "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccache"
	hpccacheaccesspolicy "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccacheaccesspolicy"
	hpccacheblobnfstarget "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccacheblobnfstarget"
	hpccacheblobtarget "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccacheblobtarget"
	hpccachenfstarget "github.com/upbound/provider-azure/internal/controller/cluster/storagecache/hpccachenfstarget"
)

// Setup_storagecache creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_storagecache(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		hpccache.Setup,
		hpccacheaccesspolicy.Setup,
		hpccacheblobnfstarget.Setup,
		hpccacheblobtarget.Setup,
		hpccachenfstarget.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_storagecache creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_storagecache(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		hpccache.SetupGated,
		hpccacheaccesspolicy.SetupGated,
		hpccacheblobnfstarget.SetupGated,
		hpccacheblobtarget.SetupGated,
		hpccachenfstarget.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
