package gif

import (
	"fmt"
	"os"

	"github.com/peterhellberg/giphy"
)

func Search(query string) []string {
	var results []string

	client := giphy.DefaultClient
	client.APIKey = "WPWGNRtoMq37sWokCR2GGIiHIWXQlPRG"
	client.Rating = "r"
	client.Limit = 50

	query_arr := []string{query}
	if res, err := client.Search(query_arr); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		for _, d := range res.Data {
			//fmt.Println(i, d.MediaURL())
			results = append(results, d.MediaURL())
		}
	}
	return results
}
