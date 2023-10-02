package semgrep

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableDeployments(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "semgrep_deployments",
		List: &plugin.ListConfig{
			Hydrate: listDeployments,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "slug", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listDeployments(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	endpoint := "/deployments"

	jsonString, err := connect(ctx, d, endpoint)

	if err != nil {
		plugin.Logger(ctx).Error("semgrep_deployments.listDeployments", "connection_error", err)
		return nil, err
	}

	var response DeploymentResponse
	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		plugin.Logger(ctx).Error("semgrep_deployments.listDeployments", "failed_unmarshal", err)
	}

	for _, deployment := range response.Deployments {
		d.StreamListItem(ctx, deployment)
	}

	return response, nil
}

//// Custom Structs

type Deployment struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Findings struct {
		URL string `json:"url"`
	} `json:"findings"`
}

type DeploymentResponse struct {
	Deployments []Deployment `json:"deployments"`
}
