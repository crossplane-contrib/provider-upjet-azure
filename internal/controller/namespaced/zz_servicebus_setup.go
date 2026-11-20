// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	namespaceauthorizationrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicebus/namespaceauthorizationrule"
	namespacedisasterrecoveryconfig "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicebus/namespacedisasterrecoveryconfig"
	queue "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicebus/queue"
	queueauthorizationrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicebus/queueauthorizationrule"
	servicebusnamespace "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicebus/servicebusnamespace"
	subscription "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicebus/subscription"
	subscriptionrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicebus/subscriptionrule"
	topic "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicebus/topic"
	topicauthorizationrule "github.com/upbound/provider-azure/v2/internal/controller/namespaced/servicebus/topicauthorizationrule"
)

// Setup_servicebus creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicebus(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		namespaceauthorizationrule.Setup,
		namespacedisasterrecoveryconfig.Setup,
		queue.Setup,
		queueauthorizationrule.Setup,
		servicebusnamespace.Setup,
		subscription.Setup,
		subscriptionrule.Setup,
		topic.Setup,
		topicauthorizationrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_servicebus creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_servicebus(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		namespaceauthorizationrule.SetupGated,
		namespacedisasterrecoveryconfig.SetupGated,
		queue.SetupGated,
		queueauthorizationrule.SetupGated,
		servicebusnamespace.SetupGated,
		subscription.SetupGated,
		subscriptionrule.SetupGated,
		topic.SetupGated,
		topicauthorizationrule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_servicebus registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_servicebus(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		namespaceauthorizationrule.SetupWebhookWithManager,
		namespacedisasterrecoveryconfig.SetupWebhookWithManager,
		queue.SetupWebhookWithManager,
		queueauthorizationrule.SetupWebhookWithManager,
		servicebusnamespace.SetupWebhookWithManager,
		subscription.SetupWebhookWithManager,
		subscriptionrule.SetupWebhookWithManager,
		topic.SetupWebhookWithManager,
		topicauthorizationrule.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
