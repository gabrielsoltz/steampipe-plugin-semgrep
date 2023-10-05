package semgrep

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData, endpoint string) (string, error) {

	var baseUrl, token string

	// Prefer config options given in Steampipe
	semgrepConfig := GetConfig(d.Connection)

	baseUrl = os.Getenv("SEMGREP_URL")
	if semgrepConfig.BaseUrl != nil {
		baseUrl = *semgrepConfig.BaseUrl
	}

	token = os.Getenv("SEMGREP_TOKEN")
	if semgrepConfig.Token != nil {
		token = *semgrepConfig.Token
	}

	if baseUrl == "" {
		return "", errors.New("'baseUrl' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if token == "" {
		return "", errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	// Create a new request
	req, err := http.NewRequest("GET", baseUrl+endpoint, nil)
	if err != nil {
		plugin.Logger(ctx).Error("Failed to create request: %v", err)
		return "", err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		plugin.Logger(ctx).Error("Failed to make request: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		plugin.Logger(ctx).Error("Failed to read response body: %v", err)
		return "", err
	}

	return string(body), nil
}
