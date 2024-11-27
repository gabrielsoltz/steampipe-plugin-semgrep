# Table: semgrep_sca_finding

Semgrep findings from the Semgrep SCA (Supply Chain) module.

## Examples

### List all SCA findings

```sql
select
  id,
  state,
  repository,
  triage_state,
  severity,
  confidence,
  rule
from
  semgrep_sca_finding;
```

### Group SCA findings by severity for repository gabrielsoltz/steampipe-plugin-semgrep

```sql
select
  count(*) as findings,
  severity
from
  semgrep_sca_finding
where
  repository ->> 'name' = 'gabrielsoltz/steampipe-plugin-semgrep'
group by
  severity;
```
