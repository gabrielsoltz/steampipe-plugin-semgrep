package main

import (
	"github.com/gabrielsoltz/steampipe-plugin-semgrep/semgrep"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: semgrep.Plugin})
}
