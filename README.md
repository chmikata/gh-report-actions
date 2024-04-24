# gh-report-actions - Register GitHub Issue for GitHub Actions

## Introduction

This action manipulates a GitHub Issue.
It allows you to register an Issue and update the Issue with the same title.

> [!WARNING]
> Not all REST APIs are supported.

## Parameter Reference

### inputs

| Name    | Type     | Required | Default | Description                                              |
| ------- | -------- | -------- | ------- | -------------------------------------------------------- |
| `org`   | `String` | `true`   |         | Name of the organization to which the repository belongs |
| `repo`  | `String` | `true`   |         | Repository name                                          |
| `token` | `String` | `true`   |         | PAT or GitHub Token                                      |
| `title` | `String` | `true`   |         | Issue title                                              |
| `input` | `String` | `true`   |         | File path of the content to be registered in the Issue   |
| `label` | `String` | `true`   |         | Label to be assigned to the issue                        |

### outputs

| Name     | Type   | Description               |
| -------- | ------ | ------------------------- |
| `result` | `Json` | Registered Issue Contents |

## Usage

```yaml
name: ci

on:
  push:
    branches:
      - 'main'

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - name: Report Issue
        id: report
        uses: chmikata/gh-report-actions@v1
        with:
          org: organization
          repo: repository
          token: ${{ secrets.PAT }}
          title: title
          input: report/body.out
          label: label
      - name: Output Result
        run: |
          echo ${{ steps.report.outputs.result }}
```
The following output results
```bash
{"id":100,"number":10,"title":"title-string","body":"body-string"}
```
