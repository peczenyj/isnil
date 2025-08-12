# isnil

[![tag](https://img.shields.io/github/tag/peczenyj/isnil.svg)](https://github.com/peczenyj/isnil/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.23-%23007d9c)
[![GoDoc](https://pkg.go.dev/badge/github.com/peczenyj/isnil)](http://pkg.go.dev/github.com/peczenyj/isnil)
[![Go](https://github.com/peczenyj/isnil/actions/workflows/go.yml/badge.svg)](https://github.com/peczenyj/isnil/actions/workflows/go.yml)
[![Lint](https://github.com/peczenyj/isnil/actions/workflows/lint.yml/badge.svg)](https://github.com/peczenyj/isnil/actions/workflows/lint.yml)
[![codecov](https://codecov.io/gh/peczenyj/isnil/graph/badge.svg?token=9y6f3vGgpr)](https://codecov.io/gh/peczenyj/isnil)
[![Report card](https://goreportcard.com/badge/github.com/peczenyj/isnil)](https://goreportcard.com/report/github.com/peczenyj/isnil)
[![CodeQL Advanced](https://github.com/peczenyj/isnil/actions/workflows/codeql.yml/badge.svg)](https://github.com/peczenyj/isnil/actions/workflows/codeql.yml)
[![Dependency Review](https://github.com/peczenyj/isnil/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/peczenyj/isnil/actions/workflows/dependency-review.yml)
[![License](https://img.shields.io/github/license/peczenyj/isnil)](./LICENSE)
[![Latest release](https://img.shields.io/github/release/peczenyj/isnil.svg)](https://github.com/peczenyj/isnil/releases/latest)
[![GitHub Release Date](https://img.shields.io/github/release-date/peczenyj/isnil.svg)](https://github.com/peczenyj/isnil/releases/latest)
[![Last commit](https://img.shields.io/github/last-commit/peczenyj/isnil.svg)](https://github.com/peczenyj/isnil/commit/HEAD)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/peczenyj/isnil/blob/main/CONTRIBUTING.md#pull-request-process)

Golang nil checker on steroids.

Inspired by [go-x-pkg/isnil](https://github.com/go-x-pkg/isnil) and [golang-utils/isnil](https://gitlab.com/golang-utils/isnil).

## Description

In golang, check if something is nil is not trivial. This packages uses `reflect` to ensure the variable contains a nil value or not.

## Example

```go
var err error = (*fs.PathError)(nil)

if err == nil {
    fmt.Println("I expected this to be true, but")
} else {
    fmt.Println("this check fails, since err != nil")
}

if isnil.IsNil(err) {
    fmt.Println("the solution is to use isnil package")
}
```
