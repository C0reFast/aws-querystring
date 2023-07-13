# aws-querystring

aws-querystring is a Go library for parse aws like URL query parameters into structs

## Install ##

```go
$ go get -u "github.com/c0refast/aws-querystring/query"
```

## Usage ##

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/c0refast/aws-querystring/query"
)

type TagRequest struct {
	Action       string   `query:"Action"`
	RegionID     string   `query:"RegionId"`
	ResourceIds  []string `query:"ResourceId"`
	ResourceType string   `query:"ResourceType"`
	Tags         []Tag    `query:"Tag"`
}

type Tag struct {
	Key   string `query:"Key"`
	Value string `query:"Value"`
}

func main() {
	queryStr := "Action=TagResources&RegionId=cn-hangzhou&ResourceId.1=i-bp1j6qtvdm8w0z1o0&ResourceId.2=i-bp1j6qtvdm8w0z1oP&ResourceType=instance&Tag.1.Key=TestKey&Tag.1.Value=TestValue&Tag.2.Key=TestKey&Tag.2.Value=TestValue"
	urlValues, _ := url.ParseQuery(queryStr)
	req := TagRequest{}
	err := query.BindQuery(urlValues, &req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonOutput, _ := json.MarshalIndent(req, "", "  ")
	fmt.Println("Unmarshaled output:", string(jsonOutput))
}
```