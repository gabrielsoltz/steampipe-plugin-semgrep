package main

import (
    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
    "github.com/gabrielsoltz/steampipe-plugin-semgrep/semgrep"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{PluginFunc: semgrep.Plugin})
}