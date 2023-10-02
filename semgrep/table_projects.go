package semgrep

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableProjects(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "semgrep_projects",
		List: &plugin.ListConfig{
			Hydrate:    listProjects,
			KeyColumns: plugin.SingleColumn("deployment_slug"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "url", Type: proto.ColumnType_STRING},
			{Name: "latest_scan", Type: proto.ColumnType_TIMESTAMP},
			{Name: "deployment_slug", Type: proto.ColumnType_STRING, Transform: transform.FromQual("deployment_slug")},
		},
	}
}

//// LIST FUNCTION

func listProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	deployment_slug := d.EqualsQuals["deployment_slug"].GetStringValue()

	endpoint := "/deployments/" + deployment_slug + "/projects"

	jsonString, err := connect(ctx, d, endpoint)

	if err != nil {
		plugin.Logger(ctx).Error("semgrep_projects.listProjects", "connection_error", err)
		return nil, err
	}

	var response ProjectsResponse
	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		plugin.Logger(ctx).Error("semgrep_projects.listProjects", "failed_unmarshal", err)
	}

	for _, project := range response.Projects {
		d.StreamListItem(ctx, project)
	}

	return response, nil
}

//// Custom Structs

type Project struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	LatestScan string `json:"latest_scan_at"`
}

type ProjectsResponse struct {
	Projects []Project `json:"projects"`
}
