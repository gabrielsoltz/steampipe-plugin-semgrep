---
organization: gabrielsoltz
category: ["security"]
icon_url: "/images/plugins/gabrielsoltz/semgrep.svg"
brand_color: "#0095E5"
display_name: "semgrep"
short_name: "semgrep"
description: "Steampipe plugin to query Semgrep"
og_description: "Query Semgrep with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/gabrielsoltz/semgrep-social-graphic.png"
---

# Semgrep + Steampipe

[Semgrep](https://semgrep.dev/) is a Fast, customizable, and developer-oriented SAST. Scan 30+ languages with 2,750+ Community and Pro rules.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Query your security findings and filter by state:

```sql
select
  id,
  state,
  rule_message
from
  semgrep_findings
where
  deployment_slug = 'my-company' and
  state = 'unresolved';
```

```
+----------+------------+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------->
| id       | state      | rule_message                                                                                                                                                                                 >
+----------+------------+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------->
| 0123 | unresolved | Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLA>
| 1230 | unresolved | Exposing host's Docker socket to containers via a volume. The owner of this socket is root. Giving someone access to it is equivalent to giving unrestricted root access to your host. Remove>
| 2301 | unresolved | An action sourced from a third-party repository on GitHub is not pinned to a full length commit SHA. Pinning an action to a full length commit SHA is currently the only way to use an action>
```

Query your Semgrep projects:

```sql
select
  id,
  name
from
  semgrep_projects
where
  deployment_slug = 'my-company';
```

```
+--------+----------------------------------------+
| id     | name                                   |
+--------+----------------------------------------+
| 261867 | gabrielsoltz/steampipe-plugin-semgrep  |
```

## Documentation

- **[Table definitions & examples](tables/index.md)**

## Get started

### Install

Download and install the latest Semgrep plugin:

```bash
steampipe plugin install semgrep
```

### Configuration

Installing the latest Semgrep plugin will create a config file (`~/.steampipe/config/semgrep.spc`) with a single connection named `semgrep`:

```hcl
connection "semgrep" {
  plugin = "gabrielsoltz/semgrep"

  base_url = "https://semgrep.dev/api/v1"

  # Access Token for which to use for the API
  token = ""
}
```

- `token` - Required access token from Semgrep

## Get involved

- Open source: https://github.com/gabrielsoltz/steampipe-plugin-semgrep
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
