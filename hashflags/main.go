package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Hashflag struct {
	CampaignName string `json:"campaignName"`
	Hashtag      string `json:"hashtag"`
	// Primary Key
	AssetURL            string `json:"assetUrl"`
	StartingTimestampMs string `json:"startingTimestampMs"`
	EndingTimestampMs   string `json:"endingTimestampMs"`
}

//go:embed data2.json
var local string

func getActiveHashFlags() ([]Hashflag, error) {
	// Original source
	// TODO: Test with 24 hour time
	// var url string

	// url = fmt.Sprintf(
	// 	"https://pbs.twimg.com/hashflag/config-%s.json",
	// 	time.Now().UTC().Format("2006-01-02-15"),
	// )

	// Make GET Request
	// resp, err := http.Get(url)
	// if err != nil {
	// 	return []Hashflag{}, err
	// }
	// // Check if request was successful
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return []Hashflag{}, err
	// }

	// Parse Response to go struct
	result := make([]Hashflag, 5000)
	json.Unmarshal([]byte(local), &result)

	return result, nil
}

func main() {
	hashflags, err := getActiveHashFlags()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get image size using HEAD Request
	for i, hashflag := range hashflags {
		if i > 30 {
			break
		}
		time.Sleep(time.Second * 2)
		//log
		fmt.Println(hashflag.AssetURL)
		// make HEAD Request
		res, err := http.Head(hashflag.AssetURL)
		if err != nil {
			fmt.Println(err)
			return
		}

		s := res.Header.Get("Content-Length")

		size, err := strconv.Atoi(s)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(size / 1000)
		fmt.Println()
	}

	// res, err := json.Marshal(h)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Print(string(res))
}
