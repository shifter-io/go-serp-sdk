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

