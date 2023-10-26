package botservice

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures authorization group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_bot_channel_sms", func(r *config.Resource) {
		r.Path = "botchannelsms"
	})
}
