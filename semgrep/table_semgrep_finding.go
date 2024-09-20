package semgrep

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableFinding(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "semgrep_finding",
		Description: "Table for querying Semgrep code or supply chain findings data.",
		List: &plugin.ListConfig{
			ParentHydrate: listDeployments,
			Hydrate:       listFindings,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "deployment_slug", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of this finding."},
			{Name: "ref", Type: proto.ColumnType_STRING, Description: "External reference to the source of this finding (e.g. PR)."},
			{Name: "first_seen_scan_id", Type: proto.ColumnType_INT, Description: "Unique ID of the Semgrep scan that first identified this finding."},
			{Name: "syntactic_id", Type: proto.ColumnType_STRING, Description: "ID calculated based on a finding's file path, rule identifier and matched code, and index."},
			{Name: "match_based_id", Type: proto.ColumnType_STRING, Description: "ID calculated based on a finding's file path, rule id, and the rule index."},
			{Name: "external_ticket", Type: proto.ColumnType_JSON, Description: "External ticket associated with finding."},
			{Name: "repository", Type: proto.ColumnType_JSON, Description: "Which repository is this finding a part of, defined via name."},
			{Name: "line_of_code_url", Type: proto.ColumnType_STRING, Description: "The source URL including file and line number.."},
			{Name: "triage_state", Type: proto.ColumnType_STRING, Description: "The finding's triage state. Set by the user and used along with state to generate the final status viewable in the UI."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "The finding's resolution state. Managed only by changes detected at scan time, the state is combined with triage_state to ultimately determine a final status which is exposed in the UI and API."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The finding's status as exposed in the UI. Status is a derived property combining information from the finding state and triage_state. The triage_state can be used to override the scan state if the finding is still detected."},
			{Name: "severity", Type: proto.ColumnType_STRING, Description: "Severity of the finding, derived from the rule that triggered it. Low is equivalent to INFO, Medium to WARNING, and High to ERROR."},
			{Name: "confidence", Type: proto.ColumnType_STRING, Description: "Confidence of the finding, derived from the rule that triggered it."},
			{Name: "categories", Type: proto.ColumnType_JSON, Description: "The categories of the finding as classified by the associated rule metdata."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when this finding was created."},
			{Name: "relevant_since", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when this finding was detected by Semgrep (the first time, or when reintroduced)."},
			{Name: "location", Type: proto.ColumnType_JSON, Description: "Location of the record in a file, as reported by Semgrep. If null, then the information does not exist or lacks integrity (older or broken scans)."},
			{Name: "sourcing_policy", Type: proto.ColumnType_JSON, Description: "Reference to a policy, with some basic information. If null, then the information does not exist or lacks integrity (older or broken scans)."},
			{Name: "triaged_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the finding was triaged."},
			{Name: "triage_comment", Type: proto.ColumnType_STRING, Description: "The detailed comment provided during triage."},
			{Name: "triage_reason", Type: proto.ColumnType_STRING, Description: "Reason provided when this issue was triaged."},
			{Name: "state_updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When this issue's state (resolution state) was last updated, as distinct from when the issue was triaged (triaged_at)."},
			{Name: "rule", Type: proto.ColumnType_JSON, Description: "Rule that applies to this finding."},
			{Name: "assistant", Type: proto.ColumnType_JSON, Description: "Semgrep Assistant data. Only present if Assistant is enabled."},
			{Name: "deployment_slug", Type: proto.ColumnType_STRING, Description: "Sanitized machine-readable name of the deployment."},
		},
	}
}

//// LIST FUNCTION

func listFindings(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	deployment := h.Item.(Deployment)
	if (d.EqualsQualString("deployment_slug") != "") && d.EqualsQualString("deployment_slug") != deployment.Slug {
		return nil, nil
	}

	endpoint := "/deployments/" + deployment.Slug + "/findings"

	emptyParams := map[string]string{}
	paginatedResponse, err := paginatedResponse(ctx, d, endpoint, emptyParams)

	if err != nil {
		plugin.Logger(ctx).Error("semgrep_finding.listFindings", "connection_error", err)
		return nil, err
	}

	for _, split_response := range paginatedResponse {
		var response FindingsResponse
		err = json.Unmarshal([]byte(split_response), &response)
		if err != nil {
			plugin.Logger(ctx).Error("semgrep_finding.listFindings", "failed_unmarshal", err)
		}

		for _, finding := range response.Findings {
			finding.DeploymentSlug = deployment.Slug
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
	ExternalTicket  Ticket     `json:"external_ticket"`
	Repository      Repository `json:"repository"`
	LineOfCodeURL   string     `json:"line_of_code_url"`
	TriageState     string     `json:"triage_state"`
	State           string     `json:"state"`
	Status          string     `json:"status"`
	Severity        string     `json:"severity"`
	Confidence      string     `json:"confidence"`
	Categories      []string   `json:"categories"`
	CreatedAt       string     `json:"created_at"`
	RelevantSince   string     `json:"relevant_since"`
	Location        Location   `json:"location"`
	SourcingPolicy  Policy     `json:"sourcing_policy"`
	TriagedAt       string     `json:"triaged_at"`
	TriageComment   string     `json:"triage_comment"`
	TriageReason    string     `json:"triage_reason"`
	StateUpdatedAt  string     `json:"state_updated_at"`
	Rule            Rule       `json:"rule"`
	Assistant       Assistant  `json:"assistant"`
	DeploymentSlug  string     `json:"-"`
}

type Ticket struct {
	ExternalSlug string `json:"external_slug"`
	URL          string `json:"url"`
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

type Rule struct {
	Name                 string   `json:"name"`
	Message              string   `json:"message"`
	Confidence           string   `json:"confidence"`
	Category             string   `json:"category"`
	SubCategories        []string `json:"subcategories"`
	VulnerabilityClasses []string `json:"vulnerability_classes"`
	CWENames             []string `json:"cwe_names"`
	OWASPNames           []string `json:"owasp_names"`
}

type Assistant struct {
	AutoFix struct {
		FixCode     string `json:"fix_code"`
		Explanation string `json:"explanation"`
	} `json:"autofix"`
	Guidance struct {
		Summary      string `json:"summary"`
		Instructions string `json:"instructions"`
	} `json:"guidance"`
	AutoTriage struct {
		Veredict string `json:"veredict"`
		Reason   string `json:"reason"`
	} `json:"autotriage"`
	Component struct {
		Tag  string `json:"tag"`
		Risk string `json:"risk"`
	} `json:"component"`
}
