[![Go Report Card](https://goreportcard.com/badge/github.com/mskrha/gosmtp)](https://goreportcard.com/report/github.com/mskrha/gosmtp)

## gosmtp

### Description
Very simple SMTP library.

### Installation
`go get github.com/mskrha/gosmtp`

### Usage
```go
package main

import (
	"fmt"

	"github.com/mskrha/gosmtp"
)

func main() {
	msg := gosmtp.Message{
		From:    "test1@local",
		To:      "test2@local",
		Subject: "Test",
		Body:    "Test",
	}

	smtp, err := gosmtp.NewServer("relay.local.domain", "My simple mail 0.1")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := smtp.Send(msg); err != nil {
		fmt.Println(err)
	}
}
```
