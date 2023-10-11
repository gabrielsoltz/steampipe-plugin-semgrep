# Table: semgrep_deployment

Semgrep deployments refer to the process of implementing and using the Semgrep tool within a software development or security workflow. Semgrep deployments involve the strategic integration of the Semgrep static code analysis tool into a software development pipeline or security assessment process.

## Examples

### List all Semgrep deployments

```sql
select
   id,
   name,
   slug
from
   semgrep_deployment;
```

### List Semgrep deployment with name `your-deployment`

```sql
select
   *
from
   semgrep_deployment
where
   name = 'you-deployment';
```
