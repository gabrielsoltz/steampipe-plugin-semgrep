# Table: semgrep_project

A Semgrep project is a focused and organized endeavor undertaken by a development or security team to leverage the capabilities of the Semgrep static code analysis tool. It involves the application of Semgrep to analyze and improve the quality, security, and compliance of one or more software projects.

## Examples

### List all Semgrep projects

```sql
select
  id,
  name,
  latest_scan
from
  semgrep_project;
```

### List all Semgrep projects for a specific deployment

```sql
select
  id,
  name,
  latest_scan
from
  semgrep_project
where
  deployment_slug = 'my-deployment';
```

### List all Semgrep projects with contains the tag `security`

```sql
select
  *
from
  semgrep_project
where
  tags ? 'security';
```

### List all Semgrep projects with a scan in the last 7 days

```sql
select
  *
from
  semgrep_project
where
  latest_scan > now() - interval '7 days';
```
