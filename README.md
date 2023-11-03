![image](https://hub.steampipe.io/images/plugins/gabrielsoltz/semgrep-social-graphic.png)

# Semgrep Plugin for Steampipe

Use SQL to query your security findings from [Semgrep](https://semgrep.dev/)

- **[Get started →](https://hub.steampipe.io/plugins/gabrielsoltz/semgrep)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/gabrielsoltz/semgrep/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/gabrielsoltz/steampipe-plugin-semgrep/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install gabrielsoltz/semgrep
```

Configure the API token in `~/.steampipe/config/semgrep.spc`:

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

Or through environment variables:

```shell
export SEMGREP_URL=https://semgrep.dev/api/v1
export SEMGREP_TOKEN=45f86adc2nv54efd76151530rr629fc8953c2a111111fd74fa7d361d70e55759
```

Run a query:

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

## Development

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/gabrielsoltz/steampipe-plugin-semgrep.git
cd steampipe-plugin-semgrep
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/semgrep.spc
```

Try it!

```
steampipe query
> .inspect semgrep
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/gabrielsoltz/steampipe-plugin-semgrep/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Semgrep Plugin](https://github.com/gabrielsoltz/steampipe-plugin-semgrep/labels/help%20wanted)
