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
		Name:        "semgrep_finding",
		Description: "Table for querying Semgrep findings data, including issues and metadata.",
		List: &plugin.ListConfig{
			Hydrate:    listFindings,
			KeyColumns: plugin.SingleColumn("deployment_slug"),
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of this finding."},
			{Name: "ref", Type: proto.ColumnType_STRING, Description: "External reference to the source of this finding (e.g. PR)."},
			{Name: "first_seen_scan_id", Type: proto.ColumnType_INT, Description: "First seen scan."},
			{Name: "syntactic_id", Type: proto.ColumnType_STRING, Description: "Syntatic id."},
			{Name: "match_based_id", Type: proto.ColumnType_STRING, Description: "ID calculated based on a finding's file path, rule id, and the rule index."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "Status of the finding's resolution."},
			{Name: "repository", Type: proto.ColumnType_JSON, Description: "Which repository is this finding a part of, defined via name."},
			{Name: "triage_state", Type: proto.ColumnType_STRING, Description: "Status of the finding's triaging."},
			{Name: "severity", Type: proto.ColumnType_STRING, Description: "Severity of the rule that triggered the finding. Ranges from low, which would correlate to info, up to high which would correlate to error."},
			{Name: "confidence", Type: proto.ColumnType_STRING, Description: "Confidence of the rule that triggered the finding."},
			{Name: "categories", Type: proto.ColumnType_JSON, Description: "The categories of the finding as classified by the associated rule metdata."},
			{Name: "relevant_since", Type: proto.ColumnType_TIMESTAMP, Description: "Relevant since."},
			{Name: "rule_name", Type: proto.ColumnType_STRING, Description: "Rule name of rule triggering.."},
			{Name: "rule_message", Type: proto.ColumnType_STRING, Description: "Rule message on the time of rule triggering. Older findings might have the value missing/removed."},
			{Name: "location", Type: proto.ColumnType_JSON, Description: "Location of the record in a file, as reported by Semgrep. If null, then the information does not exist or lacks integrity (older or broken scans)."},
			{Name: "sourcing_policy", Type: proto.ColumnType_JSON, Description: "Reference to a policy, with some basic information. If null, then the information does not exist or lacks integrity (older or broken scans)."},
			{Name: "triaged_at", Type: proto.ColumnType_TIMESTAMP, Description: "Triaged at."},
			{Name: "triage_comment", Type: proto.ColumnType_STRING, Description: "Triage comment."},
			{Name: "state_updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When this issues' state was last updated."},
			{Name: "deployment_slug", Type: proto.ColumnType_STRING, Transform: transform.FromQual("deployment_slug"), Description: "Sanitized machine-readable name of the deployment."},
		},
	}
}

//// LIST FUNCTION

func listFindings(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	deployment_slug := d.EqualsQualString("deployment_slug")

	endpoint := "/deployments/" + deployment_slug + "/findings"

	paginatedResponse, err := paginatedResponse(ctx, d, endpoint)

	if err != nil {
		plugin.Logger(ctx).Error("semgrep_projects.listFindings", "connection_error", err)
		return nil, err
	}

	for _, split_response := range paginatedResponse {
		var response FindingsResponse
		err = json.Unmarshal([]byte(split_response), &response)
		if err != nil {
			plugin.Logger(ctx).Error("semgrep_projects.listFindings", "failed_unmarshal", err)
		}

		for _, finding := range response.Findings {
			d.StreamListItem(ctx, finding)
		}
	}

	return paginatedResponse, nil
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
