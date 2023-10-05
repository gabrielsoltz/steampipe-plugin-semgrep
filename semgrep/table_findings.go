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
			{Name: "ref", Type: proto.ColumnType_STRING},
			{Name: "first_seen_scan_id", Type: proto.ColumnType_INT},
			{Name: "syntactic_id", Type: proto.ColumnType_STRING},
			{Name: "match_based_id", Type: proto.ColumnType_STRING},
			{Name: "state", Type: proto.ColumnType_STRING},
			{Name: "repository", Type: proto.ColumnType_JSON},
			{Name: "triage_state", Type: proto.ColumnType_STRING},
			{Name: "severity", Type: proto.ColumnType_STRING},
			{Name: "confidence", Type: proto.ColumnType_STRING},
			{Name: "categories", Type: proto.ColumnType_JSON},
			{Name: "relevant_since", Type: proto.ColumnType_TIMESTAMP},
			{Name: "rule_name", Type: proto.ColumnType_STRING},
			{Name: "rule_message", Type: proto.ColumnType_STRING},
			{Name: "location", Type: proto.ColumnType_JSON},
			{Name: "sourcing_policy", Type: proto.ColumnType_JSON},
			{Name: "triaged_at", Type: proto.ColumnType_TIMESTAMP},
			{Name: "triage_comment", Type: proto.ColumnType_STRING},
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
	ID              int        `json:"id"`
	Ref             string     `json:"ref"`
	FirstSeenScanID int        `json:"first_seen_scan_id"`
	SyntacticID     string     `json:"syntactic_id"`
	MatchBasedID    string     `json:"match_based_id"`
	State           string     `json:"state"`
	Repository      Repository `json:"repository"`
	TriageState     string     `json:"triage_state"`
	Severity        string     `json:"severity"`
	Confidence      string     `json:"confidence"`
	Categories      []string   `json:"categories"`
	RelevantSince   string     `json:"relevant_since"`
	RuleName        string     `json:"rule_name"`
	RuleMessage     string     `json:"rule_message"`
	Location        Location   `json:"location"`
	SourcingPolicy  Policy     `json:"sourcing_policy"`
	TriagedAt       string     `json:"triaged_at"`
	TriageComment   string     `json:"triage_comment"`
	StateUpdatedAt  string     `json:"state_updated_at"`
}

type Repository struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Location struct {
	FilePath  string `json:"file_path"`
	Line      int    `json:"line"`
	Column    int    `json:"column"`
	EndLine   int    `json:"end_line"`
	EndColumn int    `json:"end_column"`
}

type Policy struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type FindingsResponse struct {
	Findings []Finding `json:"findings"`
}
