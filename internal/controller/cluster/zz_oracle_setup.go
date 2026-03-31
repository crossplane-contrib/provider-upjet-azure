// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	autonomousdatabase "github.com/upbound/provider-azure/v2/internal/controller/cluster/oracle/autonomousdatabase"
	autonomousdatabasebackup "github.com/upbound/provider-azure/v2/internal/controller/cluster/oracle/autonomousdatabasebackup"
	autonomousdatabaseclonefrombackup "github.com/upbound/provider-azure/v2/internal/controller/cluster/oracle/autonomousdatabaseclonefrombackup"
	autonomousdatabaseclonefromdatabase "github.com/upbound/provider-azure/v2/internal/controller/cluster/oracle/autonomousdatabaseclonefromdatabase"
)

// Setup_oracle creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_oracle(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autonomousdatabase.Setup,
		autonomousdatabasebackup.Setup,
		autonomousdatabaseclonefrombackup.Setup,
		autonomousdatabaseclonefromdatabase.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_oracle creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_oracle(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autonomousdatabase.SetupGated,
		autonomousdatabasebackup.SetupGated,
		autonomousdatabaseclonefrombackup.SetupGated,
		autonomousdatabaseclonefromdatabase.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
