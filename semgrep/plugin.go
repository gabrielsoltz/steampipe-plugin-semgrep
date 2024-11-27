package semgrep

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-semgrep",
		DefaultTransform: transform.FromGo().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"semgrep_deployment":  tableDeployment(ctx),
			"semgrep_finding":     tableFinding(ctx),
			"semgrep_project":     tableProject(ctx),
			"semgrep_sca_finding": tableScaFinding(ctx),
		},
	}
	return p
}
