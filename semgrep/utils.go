package semgrep

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

func connect(_ context.Context, d *plugin.QueryData) (*semgrep.Client, error) {

	// Start with an empty Turbot config
	tokenProvider := semgrep.BasicAuthTransport{}
	var baseUrl, token string

	// Prefer config options given in Steampipe
	semgrepConfig := GetConfig(d.Connection)

	if semgrepConfig.BaseUrl != nil {
		baseUrl = *semgrepConfig.BaseUrl
	}
	if semgrepConfig.Token != nil {
		token = *semgrepConfig.Token
	}

	if baseUrl == "" {
		return nil, errors.New("'base_url' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}
	tokenProvider.Password = token

	// Create the client
	client, err := semgrep.NewClient(tokenProvider.Client(), baseUrl)
	if err != nil {
		return nil, fmt.Errorf("error creating Semgrep client: %s", err.Error())
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	// Done
	return client, nil
}