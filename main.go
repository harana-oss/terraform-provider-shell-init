package main

import (
	"github.com/harana-oss/terraform-provider-shell-init/shell"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: shell.Provider})
}
