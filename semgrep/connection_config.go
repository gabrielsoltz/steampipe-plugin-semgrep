package semgrep

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type semgrepConfig struct {
	BaseUrl  *string `cty:"base_url"`
	Token    *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"base_url": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &semgrepConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) semgrepConfig {
	if connection == nil || connection.Config == nil {
		return semgrepConfig{}
	}
	config, _ := connection.Config.(semgrepConfig)
	return config
}
