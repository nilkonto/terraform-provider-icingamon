package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-icingamon/icingamon"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: icingamon.Provider})
}
