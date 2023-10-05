# Semgrep Plugin for Steampipe

Use SQL to query your security findings from Semgrep.

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install gabrielsoltz/semgrep
```

Run a query:

```sql
select id, state, rule_message from semgrep_findings where state = 'unresolved' and deployment_slug = 'my-company';
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
