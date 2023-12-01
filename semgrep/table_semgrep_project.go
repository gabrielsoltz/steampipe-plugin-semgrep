package semgrep

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableProject(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "semgrep_project",
		Description: "Table for querying Semgrep projects, containing project-specific information and configurations.",
		List: &plugin.ListConfig{
			ParentHydrate: listDeployments,
			Hydrate:       listProjects,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "deployment_slug", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of this project."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of this project."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of this project."},
			{Name: "latest_scan", Type: proto.ColumnType_TIMESTAMP, Description: "Latest scan date of this project."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags of this project."},
			{Name: "deployment_slug", Type: proto.ColumnType_STRING, Transform: transform.FromQual("deployment_slug"), Description: "Sanitized machine-readable name of the deployment."},
		},
	}
}

//// LIST FUNCTION

func listProjects(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	deployment := h.Item.(Deployment)
	if (d.EqualsQualString("deployment_slug") != "") && d.EqualsQualString("deployment_slug") != deployment.Slug {
		return nil, nil
	}

	endpoint := "/deployments/" + deployment.Slug + "/projects"

	paginatedResponse, err := paginatedResponse(ctx, d, endpoint)

	if err != nil {
		plugin.Logger(ctx).Error("semgrep_project.listProjects", "connection_error", err)
		return nil, err
	}

	for _, split_response := range paginatedResponse {
		var response ProjectsResponse
		err = json.Unmarshal([]byte(split_response), &response)
		if err != nil {
			plugin.Logger(ctx).Error("semgrep_project.listProjects", "failed_unmarshal", err)
		}

		for _, project := range response.Projects {
			d.StreamListItem(ctx, project)
		}
	}

	return paginatedResponse, nil
}

//// Custom Structs

type Project struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Url        string   `json:"url"`
	LatestScan string   `json:"latest_scan_at"`
	Tags       []string `json:"tags"`
}

type ProjectsResponse struct {
	Projects []Project `json:"projects"`
}
