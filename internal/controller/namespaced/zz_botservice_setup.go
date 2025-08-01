// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	botchannelalexa "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelalexa"
	botchanneldirectline "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchanneldirectline"
	botchannelline "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelline"
	botchannelmsteams "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelmsteams"
	botchannelslack "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelslack"
	botchannelsms "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelsms"
	botchannelsregistration "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelsregistration"
	botchannelwebchat "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botchannelwebchat"
	botconnection "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botconnection"
	botwebapp "github.com/upbound/provider-azure/internal/controller/namespaced/botservice/botwebapp"
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

// SetupGated_botservice creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_botservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		botchannelalexa.SetupGated,
		botchanneldirectline.SetupGated,
		botchannelline.SetupGated,
		botchannelmsteams.SetupGated,
		botchannelslack.SetupGated,
		botchannelsms.SetupGated,
		botchannelsregistration.SetupGated,
		botchannelwebchat.SetupGated,
		botconnection.SetupGated,
		botwebapp.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
