package semgrep

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableScaFinding(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "semgrep_sca_finding",
		Description: "Table for querying Semgrep supply chain findings data.",
		List: &plugin.ListConfig{
			ParentHydrate: listDeployments,
			Hydrate:       listScaFindings,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "deployment_slug", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of this finding."},
			{Name: "ref", Type: proto.ColumnType_STRING, Description: "External reference to the source of this finding (e.g. PR)."},
			{Name: "external_ticket", Type: proto.ColumnType_JSON, Description: "External ticket associated with finding."},
			{Name: "first_seen_scan_id", Type: proto.ColumnType_INT, Description: "Unique ID of the Semgrep scan that first identified this finding."},
			{Name: "syntactic_id", Type: proto.ColumnType_STRING, Description: "ID calculated based on a finding's file path, rule identifier and matched code, and index."},
			{Name: "match_based_id", Type: proto.ColumnType_STRING, Description: "ID calculated based on a finding's file path, rule id, and the rule index."},
			{Name: "repository", Type: proto.ColumnType_JSON, Description: "Which repository is this finding a part of, defined via name."},
			{Name: "line_of_code_url", Type: proto.ColumnType_STRING, Description: "The source URL including file and line number.."},
			{Name: "triage_state", Type: proto.ColumnType_STRING, Description: "The finding's triage state. Set by the user and used along with state to generate the final status viewable in the UI."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "The finding's resolution state. Managed only by changes detected at scan time, the state is combined with triage_state to ultimately determine a final status which is exposed in the UI and API."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The finding's status as exposed in the UI. Status is a derived property combining information from the finding state and triage_state. The triage_state can be used to override the scan state if the finding is still detected."},
			{Name: "confidence", Type: proto.ColumnType_STRING, Description: "Confidence of the finding, derived from the rule that triggered it."},
			{Name: "categories", Type: proto.ColumnType_JSON, Description: "The categories of the finding as classified by the associated rule metdata."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when this finding was created."},
			{Name: "relevant_since", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when this finding was detected by Semgrep (the first time, or when reintroduced)."},
			{Name: "location", Type: proto.ColumnType_JSON, Description: "Location of the record in a file, as reported by Semgrep. If null, then the information does not exist or lacks integrity (older or broken scans)."},
			{Name: "triaged_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the finding was triaged."},
			{Name: "triage_comment", Type: proto.ColumnType_STRING, Description: "The detailed comment provided during triage."},
			{Name: "triage_reason", Type: proto.ColumnType_STRING, Description: "Reason provided when this issue was triaged."},
			{Name: "state_updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When this issue's state (resolution state) was last updated, as distinct from when the issue was triaged (triaged_at)."},
			{Name: "rule", Type: proto.ColumnType_JSON, Description: "Rule that applies to this finding."},
			{Name: "severity", Type: proto.ColumnType_STRING, Description: "Severity of the finding, derived from the rule that triggered it. Low is equivalent to INFO, Medium to WARNING, and High to ERROR."},
			{Name: "vulnerability_identifier", Type: proto.ColumnType_STRING, Description: "Identifier of the vulnerability that this finding is associated with."},
			{Name: "epss_score", Type: proto.ColumnType_JSON, Description: "Expected Probability and Severity score of the finding."},
			{Name: "reachability", Type: proto.ColumnType_STRING, Description: "Reachability of the finding."},
			{Name: "reachable_conditions", Type: proto.ColumnType_STRING, Description: "Reachability conditions of the finding."},
			{Name: "found_dependency", Type: proto.ColumnType_JSON, Description: "The dependency that was found to be vulnerable."},
			{Name: "usage", Type: proto.ColumnType_JSON, Description: "Usage of the dependency that was found to be vulnerable."},
			{Name: "fix_recommendation", Type: proto.ColumnType_JSON, Description: "Fix recommendation for the finding."},
			{Name: "deployment_slug", Type: proto.ColumnType_STRING, Description: "Sanitized machine-readable name of the deployment."},
		},
	}
}

//// LIST FUNCTION

func listScaFindings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	deployment := h.Item.(Deployment)
	if (d.EqualsQualString("deployment_slug") != "") && d.EqualsQualString("deployment_slug") != deployment.Slug {
		return nil, nil
	}

	endpoint := "/deployments/" + deployment.Slug + "/findings"

	emptyParams := map[string]string{
		"issue_type": "sca",
	}
	paginatedResponse, err := paginatedResponse(ctx, d, endpoint, emptyParams)

	if err != nil {
		plugin.Logger(ctx).Error("semgrep_finding.listScaFindings", "connection_error", err)
		return nil, err
	}

	for _, split_response := range paginatedResponse {
		var response ScaFindingsResponse
		err = json.Unmarshal([]byte(split_response), &response)
		if err != nil {
			plugin.Logger(ctx).Error("semgrep_finding.listScaFindings", "failed_unmarshal", err)
		}

		for _, finding := range response.Findings {
			finding.DeploymentSlug = deployment.Slug
			d.StreamListItem(ctx, finding)
		}
	}

	return paginatedResponse, nil
}

//// Custom Structs

type ScaFinding struct {
	ID                      int               `json:"id"`
	Ref                     string            `json:"ref"`
	ExternalTicket          Ticket            `json:"external_ticket"`
	FirstSeenScanID         int               `json:"first_seen_scan_id"`
	SyntacticID             string            `json:"syntactic_id"`
	MatchBasedID            string            `json:"match_based_id"`
	Repository              Repository        `json:"repository"`
	LineOfCodeURL           string            `json:"line_of_code_url"`
	TriageState             string            `json:"triage_state"`
	State                   string            `json:"state"`
	Status                  string            `json:"status"`
	Confidence              string            `json:"confidence"`
	Categories              []string          `json:"categories"`
	CreatedAt               string            `json:"created_at"`
	RelevantSince           string            `json:"relevant_since"`
	Location                Location          `json:"location"`
	TriagedAt               string            `json:"triaged_at"`
	TriageComment           string            `json:"triage_comment"`
	TriageReason            string            `json:"triage_reason"`
	StateUpdatedAt          string            `json:"state_updated_at"`
	Rule                    Rule              `json:"rule"`
	Severity                string            `json:"severity"`
	VulnerabilityIdentifier string            `json:"vulnerability_identifier"`
	EpssScore               EpssScore         `json:"epss_score"`
	Reachability            string            `json:"reachability"`
	ReachableConditions     string            `json:"reachable_conditions"`
	FoundDependency         FoundDependency   `json:"found_dependency"`
	Usage                   Usage             `json:"usage"`
	FixRecommendation       FixRecommendation `json:"fix_recommendation"`
	DeploymentSlug          string            `json:"-"`
}

type ScaFindingsResponse struct {
	Findings []ScaFinding `json:"findings"`
}

type EpssScore struct {
	Score      float32 `json:"score"`
	Percentile float32 `json:"percentile"`
}

type FoundDependency struct {
	Package         string `json:"package"`
	Version         string `json:"version"`
	Ecosystem       string `json:"ecosystem"`
	Transitivity    string `json:"transitivity"`
	LockfileLineURL string `json:"lockfile_line_url"`
}

type Usage struct {
	Location       Location `json:"location"`
	ExternalTicket Ticket   `json:"usage"`
}

type FixRecommendation struct {
	Package string `json:"package"`
	Version string `json:"version"`
}
