# http-request

[![Build Status](https://travis-ci.org/viknesh-nm/http-request.svg?branch=master)](https://travis-ci.org/viknesh-nm/http-request)

http-request is a simple package written in go for getting the http request from the URL
## Installation

```go 
go get github.com/viknesh-nm/http-request
```

## Example

```go
package main

import (
	"fmt"

	request "github.com/viknesh-nm/http-request"
)

type SampleTest struct {
	// Add the fields available in your required URL
}

func main() {
	var test SampleTest
	errG := request.DoGetHTTPRequest("<YOUR GET Request URL>", &test)
	if errG != nil {
		fmt.Println(errG)
		return
	}
	fmt.Println(test)

	var test1 SampleTest
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := request.DoHTTPRequest(
		"<METHOD>",
		"<YOUR Request URL>",
		headers,
		strings.NewReader(""),
		&test1,
	)

	fmt.Println(test1)
}
```