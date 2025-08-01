// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appactioncustom "github.com/upbound/provider-azure/internal/controller/cluster/logic/appactioncustom"
	appactionhttp "github.com/upbound/provider-azure/internal/controller/cluster/logic/appactionhttp"
	appintegrationaccount "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccount"
	appintegrationaccountbatchconfiguration "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountbatchconfiguration"
	appintegrationaccountpartner "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountpartner"
	appintegrationaccountschema "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountschema"
	appintegrationaccountsession "github.com/upbound/provider-azure/internal/controller/cluster/logic/appintegrationaccountsession"
	apptriggercustom "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggercustom"
	apptriggerhttprequest "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggerhttprequest"
	apptriggerrecurrence "github.com/upbound/provider-azure/internal/controller/cluster/logic/apptriggerrecurrence"
	appworkflow "github.com/upbound/provider-azure/internal/controller/cluster/logic/appworkflow"
)

// Setup_logic creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_logic(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appactioncustom.Setup,
		appactionhttp.Setup,
		appintegrationaccount.Setup,
		appintegrationaccountbatchconfiguration.Setup,
		appintegrationaccountpartner.Setup,
		appintegrationaccountschema.Setup,
		appintegrationaccountsession.Setup,
		apptriggercustom.Setup,
		apptriggerhttprequest.Setup,
		apptriggerrecurrence.Setup,
		appworkflow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
