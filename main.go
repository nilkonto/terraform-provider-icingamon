package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/nilkonto/terraform-provider-icingamon/icingamon"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: icingamon.Provider})
}
