package semgrep

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableFindings(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "semgrep_findings",
		List: &plugin.ListConfig{
			Hydrate:    listFindings,
			KeyColumns: plugin.SingleColumn("deployment_slug"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "state", Type: proto.ColumnType_STRING},
			{Name: "rule_name", Type: proto.ColumnType_STRING},
			{Name: "rule_message", Type: proto.ColumnType_STRING},
			{Name: "state_updated_at", Type: proto.ColumnType_TIMESTAMP},
			{Name: "deployment_slug", Type: proto.ColumnType_STRING, Transform: transform.FromQual("deployment_slug")},
		},
	}
}

//// LIST FUNCTION

func listFindings(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	deployment_slug := d.EqualsQuals["deployment_slug"].GetStringValue()

	endpoint := "/deployments/" + deployment_slug + "/findings"

	jsonString, err := connect(ctx, d, endpoint)

	if err != nil {
		plugin.Logger(ctx).Error("semgrep_findings.listFindings", "connection_error", err)
		return nil, err
	}

	var response FindingsResponse
	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		plugin.Logger(ctx).Error("semgrep_findings.listFindings", "failed_unmarshal", err)
	}

	for _, finding := range response.Findings {
		d.StreamListItem(ctx, finding)
	}

	return response, nil
}

//// Custom Structs

type Finding struct {
	ID           int    `json:"id"`
	State        string `json:"state"`
	RuleName     string `json:"rule_name"`
	RuleMessage  string `json:"rule_message"`
	StateUpdated string `json:"state_updated_at"`
}

type FindingsResponse struct {
	Findings []Finding `json:"findings"`
}
