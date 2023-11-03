---
organization: gabrielsoltz
category: ["security"]
icon_url: "/images/plugins/gabrielsoltz/semgrep.svg"
brand_color: "#13BF95"
display_name: "Semgrep"
short_name: "semgrep"
description: "Steampipe plugin to query deployments, findings, and projects from Semgrep."
og_description: "Query Semgrep with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/gabrielsoltz/semgrep-social-graphic.png"
---

# Semgrep + Steampipe

[Semgrep](https://semgrep.dev/) is a Fast, customizable, and developer-oriented SAST. Scan 30+ languages with 2,750+ Community and Pro rules.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Query your security findings and filter by state:

```sql
select
  triage_state,
  severity,
  state,
  rule_message,
  repository ->> 'name' as repo_name
from
  semgrep_finding
where
  deployment_slug = 'my-company'
  and state = 'unresolved';
```

```
+--------------+----------+------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+----------------------------------------+
| triage_state | severity | state      | rule_message                                                                                                                                                                                                                              | repo_name                              |
+--------------+----------+------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+----------------------------------------+
| untriaged    | medium   | unresolved | Detected possible formatted SQL query. Use parameterized queries instead.                                                                                                                                                                 | gabrielsoltz/steampipe-plugin-semgrep |
| untriaged    | medium   | unresolved | Service 'localstack' allows for privilege escalation via setuid or setgid binaries. Add 'no-new-privileges:true' in 'security_opt' to prevent this.                                                                                       | gabrielsoltz/steampipe-plugin-semgrep |
+--------------+----------+------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+----------------------------------------+
```

## Documentation

**[Table definitions & examples →](/plugins/gabrielsoltz/semgrep/tables)**

## Get started

### Install

Download and install the latest Semgrep plugin:

```bash
steampipe plugin install gabrielsoltz/semgrep
```

### Configuration

Installing the latest Semgrep plugin will create a config file (`~/.steampipe/config/semgrep.spc`) with a single connection named `semgrep`:

```hcl
connection "semgrep" {
  plugin = "gabrielsoltz/semgrep"

  # The base URL of Semgrep. Required.
  # This can be set via the `SEMGREP_URL` environment variable.
  # base_url = "https://semgrep.dev/api/v1"

  # The access token required for API calls. Required.
  # This can also be set via the `SEMGREP_TOKEN` environment variable.
  # token = "45f86adc2nv54efd76151530rr629fc8953c2a111111fd74fa7d361d70e55759"
}
```

- `token` - Required access token from Semgrep

Alternatively, you can also use the standard Semgrep environment variables to obtain credentials only if other arguments (base_urland token) are not specified in the connection:

```
export SEMGREP_URL=https://semgrep.dev/api/v1
export SEMGREP_TOKEN=45f86adc2nv54efd76151530rr629fc8953c2a111111fd74fa7d361d70e55759
```

## Get involved

- Open source: https://github.com/gabrielsoltz/steampipe-plugin-semgrep
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
