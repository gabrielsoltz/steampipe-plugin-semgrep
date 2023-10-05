# Table: semgrep_projects

List all your Semgrep projects. You must specify which deployment slug in the where or join clause using the `deployment_slug` column.

## Examples

### List all Semgrep projects

```sql
select
 id,
 name,
 latest_scan
from
  semgrep_projects
where
  deployment_slug = 'my-deployment';
```
