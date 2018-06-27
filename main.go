package main

import (
	"github.com/erikvanbrakel/terraform-provider-victorops/provider"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
