# gh-report-cli - CLI for GitHub Issues

## Introduction

CLI tool to manipulate the REST API of github issues.

> [!WARNING]
> Not all REST APIs are supported.

## Parameter Reference

### Command Usage

```bash
Usage:
  gh-report-cli [command] [flags]

Available Commands:
  package     Display package
  tag         Display container image tags

Flags:
  -h, --help   help for package
  ・・・

Global Flags:
  -m, --matcher string   Name of the container image to match
  -o, --org string       Organization name
  -t, --token string     Token for authentication
```

### Global Flags

| Name    | Shortened Name | Type     | Required | Default | Description                       |
| ------- | -------------- | -------- | -------- | ------- | --------------------------------- |
| `org`   | `o`            | `String` | `true`   |         | Organization name                 |
| `repo`  | `r`            | `String` | `true`   |         | Repository name                   |
| `token` | `t`            | `String` | `true`   |         | PAT or GitHub Token               |
| `title` | `T`            | `String` | `true`   |         | Issue title                       |
| `input` | `i`            | `String` | `true`   |         | Issue text                        |
| `label` | `l`            | `String` | `true`   |         | Label to be assigned to the issue |

## Command Output

### Command [report]

```bash
$ gh-report-cli report --org org --repo repo --token ********** --input input.json --label label
{"id":1,"number":5,"title":"title string","body":"body string"}
```
