package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/brewer35/terraform-provider-bitbucket/bitbucket"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bitbucket.Provider})
}
