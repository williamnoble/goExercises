package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Story describes an individual HN story/article for which users may comment.

type Story struct {
	Time        int
	Descendants int
	By          string
	Kids        []int
	Score       int
	Title       string
	Type        string
	URL         string
}

// Entry point for 'gohn'
func main() {
	data := getStoryData()
	displayStories(data)
}

const (
	uriTopStories = "https://hacker-news.firebaseio.com/v0/topstories.json"
)

// Retrieve the top 10 stories listed on Hackernews.
func getTop10() []int {
	var top10 []int
	resp, err := http.Get(uriTopStories)
	if err != nil {
		fmt.Println(err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	b, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &top10)
	if err != nil {
		fmt.Println(err)
	}
	return top10[:10]

	// Retrieve story data, return a list of stories.
}
func getStoryData() []Story {
	story := Story{}
	var stories []Story
	topStories := getTop10()

	for _, value := range topStories {
		id := strconv.Itoa(value)
		query := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%s.json", id)
		resp, err := http.Get(query)
		if err != nil {
			fmt.Println(err)
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		b, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(b, &story)

		if err != nil {
			fmt.Println(err)
		}
		stories = append(stories, story)
	}
	return stories
}

func convertUnixtime(unixTime string) string {
	// convert unixTime to int64
	unixTimeInt64, err := strconv.ParseInt(unixTime, 10, 64)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tm := time.Unix(unixTimeInt64, 0)
	timeNow := time.Now()
	hnhour := time.Time.Hour(tm)
	nowHour := time.Time.Hour(timeNow)
	hoursPastInterger := (nowHour - hnhour) - 2

	hoursPast := strconv.Itoa(hoursPastInterger)

	return hoursPast
}

func displayStories(stories []Story) {
	var o string

	for _, story := range stories {
		stringUnixtime := strconv.Itoa(story.Time)
		currentTime := convertUnixtime(stringUnixtime)
		o += fmt.Sprintf("%s\n%s\nBy %s (%s hours ago)\n\n", story.Title, story.URL, story.By, currentTime)
	}
	fmt.Println(o)

	// cmd := exec.Command("/usr/bin/less")
	// cmd.Stdin = strings.NewReader(o)
	// cmd.Stdout = os.Stdout

	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
