package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"steampipe-plugin-sggov/sggov"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: sggov.Plugin,
	})
}
