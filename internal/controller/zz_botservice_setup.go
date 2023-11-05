// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	botchannelalexa "github.com/upbound/provider-azure/internal/controller/botservice/botchannelalexa"
	botchanneldirectline "github.com/upbound/provider-azure/internal/controller/botservice/botchanneldirectline"
	botchannelline "github.com/upbound/provider-azure/internal/controller/botservice/botchannelline"
	botchannelmsteams "github.com/upbound/provider-azure/internal/controller/botservice/botchannelmsteams"
	botchannelslack "github.com/upbound/provider-azure/internal/controller/botservice/botchannelslack"
	botchannelsms "github.com/upbound/provider-azure/internal/controller/botservice/botchannelsms"
	botchannelsregistration "github.com/upbound/provider-azure/internal/controller/botservice/botchannelsregistration"
	botchannelwebchat "github.com/upbound/provider-azure/internal/controller/botservice/botchannelwebchat"
	botconnection "github.com/upbound/provider-azure/internal/controller/botservice/botconnection"
	botwebapp "github.com/upbound/provider-azure/internal/controller/botservice/botwebapp"
)

// Setup_botservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_botservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		botchannelalexa.Setup,
		botchanneldirectline.Setup,
		botchannelline.Setup,
		botchannelmsteams.Setup,
		botchannelslack.Setup,
		botchannelsms.Setup,
		botchannelsregistration.Setup,
		botchannelwebchat.Setup,
		botconnection.Setup,
		botwebapp.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
