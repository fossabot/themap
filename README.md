# TheMAP

[![GitHub Actions](https://github.com/sendtips/themap/workflows/Go/badge.svg)](https://github.com/sendtips/themap/actions?workflow=Go)
[![GoDoc](https://godoc.org/github.com/sendtips/themap?status.svg)](https://godoc.org/github.com/sendtips/themap)
[![codecov](https://codecov.io/gh/sendtips/themap/branch/master/graph/badge.svg)](https://codecov.io/gh/sendtips/themap)
[![Go Report Card](https://goreportcard.com/badge/github.com/sendtips/themap)](https://goreportcard.com/report/github.com/sendtips/themap)
[![Sourcegraph](https://sourcegraph.com/github.com/sendtips/themap/-/badge.svg)](https://sourcegraph.com/github.com/sendtips/themap?badge)

A Go library to work with [TheMAP](https://doc.mapcard.pro/).

## Install
Install by import `github.com/sendtips/themap` or via `go get github.com/sendtips/themap`

## Example
To obtain a session you need call Init method.

```go
package main

import (
	"fmt"
	"github.com/sendtips/themap"
)

func main() {
	p := themap.New("SendtipsTestTerminal", "SendtipsTestTerminal")
	//p.SetAuthUser("login", "password")
	p.SetMerch("SendtipsTestTerminal", "123")
	//p.SetTerm("123")

	err := p.Init(200) // Create session for 2.00RUB
	if err != nil {
		fmt.Printf("Error ocurred: %v", err)
	}

	fmt.Printf("%+v", p.Reply.SessionGUID) // Will have a theMAP session identifier
}
```
