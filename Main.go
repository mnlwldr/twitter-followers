package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"strconv"
)

func initAnaconda() *anaconda.TwitterApi {
	return anaconda.NewTwitterApiWithCredentials(
		"",
		"",
		"",
		"")
}

func main() {

	const count = "200"
	const skipStatus = "true"
	const includeUserEntities = "false"

	api := initAnaconda()

	params := url.Values{}
	params.Set("count", count)
	params.Set("skip_status", skipStatus)
	params.Set("include_user_entities", includeUserEntities)

	var cursor int64 = -1 // Initial value, which is the first page
	for cursor != 0 {

		params.Set("cursor", strconv.FormatInt(cursor, 10))
		followers, err := api.GetFollowersList(params)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		for _, follower := range followers.Users {
			fmt.Println(follower.ScreenName)
		}

		cursor = followers.Next_cursor
	}
}
