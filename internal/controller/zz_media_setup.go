/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	asset "github.com/upbound/provider-azure/internal/controller/media/asset"
	assetfilter "github.com/upbound/provider-azure/internal/controller/media/assetfilter"
	contentkeypolicy "github.com/upbound/provider-azure/internal/controller/media/contentkeypolicy"
	job "github.com/upbound/provider-azure/internal/controller/media/job"
	liveevent "github.com/upbound/provider-azure/internal/controller/media/liveevent"
	liveeventoutput "github.com/upbound/provider-azure/internal/controller/media/liveeventoutput"
	servicesaccount "github.com/upbound/provider-azure/internal/controller/media/servicesaccount"
	servicesaccountfilter "github.com/upbound/provider-azure/internal/controller/media/servicesaccountfilter"
	streamingendpoint "github.com/upbound/provider-azure/internal/controller/media/streamingendpoint"
	streaminglocator "github.com/upbound/provider-azure/internal/controller/media/streaminglocator"
	streamingpolicy "github.com/upbound/provider-azure/internal/controller/media/streamingpolicy"
	transform "github.com/upbound/provider-azure/internal/controller/media/transform"
)

// Setup_media creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_media(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		asset.Setup,
		assetfilter.Setup,
		contentkeypolicy.Setup,
		job.Setup,
		liveevent.Setup,
		liveeventoutput.Setup,
		servicesaccount.Setup,
		servicesaccountfilter.Setup,
		streamingendpoint.Setup,
		streaminglocator.Setup,
		streamingpolicy.Setup,
		transform.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
