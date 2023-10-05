# Table: semgrep_findings

List all your Semgrep findings. You must specify which deployment slug in the where or join clause using the `deployment_slug`` column.

## Examples

### List all Semgrep findings

```sql
select
 id,
 state,
 repository,
 triage_state,
 severity,
 confidence,
 rule_name
from
  semgrep_findings
where
  deployment_slug = 'my-deployment';
```

### List all Semgrep with high severity that are not triaged

```sql
select
 id,
 state,
 repository,
 triage_state,
 severity,
 confidence,
 rule_name
from
  semgrep_findings
where
  deployment_slug = 'my-deployment'
  and severity = 'high'
  and triage_state = 'untriaged';
```
