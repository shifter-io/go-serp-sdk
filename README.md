# Shifter's SERP Go SDK

[Shifter SERP](https://shifter.io/services/serp-scraping) is an API that allows scraping various search engines such as Google, Bing, Yandex, etc. while using rotating proxies to prevent bans. This SDK for Go makes the usage of the API easier to implement in any project you have.

## Installation

You must have Go 1.16 installed.

Run the following command to add the package:

```
go get -u github.com/shifter-io/go-serp-sdk
```

## Quick Start

Use the following code to test the SDK:

```
package main

import (
	"fmt"
	serp "github.com/shifter-io/go-serp-sdk"
)

func main()  {

	response, err := serp.GoogleSearch(serp.GoogleSearchParams{
		ApiKey: "{API_KEY}",
		Q: "pizza",
	})

	if err != nil {
		panic(err)
	}

	results := response["organic_results"].([]interface{})
	firstResult := results[0].(map[string]interface{})
	fmt.Println(firstResult["title"].(string))

}
```
