package plugins

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-template/template"
)

func init() {
	templateProvider := func() {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: template.Provider,
		})
	}
	KnownPlugins["terraform-provider-template"] = templateProvider
}
