// This is the start of my main document. Transposed from the copy to the RHS.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Story describes a single post on Hackernews
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

func main() {
	data := getStoryData()
	displayStories(data)
}

const (
	uriTopStories = "super-secret-url"
)

func getTop10() []int {
	var top10 []int
	resp, err := http.Get(uriTopStories)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &top10)
	if err != nil {
		fmt.Println(err)
	}
	return top10[:10]
}

func getStoryData() []Story {
	story := Story{}
	var stories []Story
	topStories := getTop10()

	for _, value := range topStories {
		id := strconv.Itoa(value)
		query := fmt.Sprintf("URL %s", id)
		resp, err := http.Get(query)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(b, &story)

		if err != nil {
			fmt.Println(err)
		}

		stories = append(stories, story)
	}
	return stories
}

func displayStories(stories []Story) {
	var output string
	for _, story := range stories {
		stringUnixTime := strconv.Itoa(story.Time)
		currentTime := convertUnixTime(stringUnixTime)
		fmt.Println(output, currentTime)
	}
}

func convertUnixTime(stringUnixTime string) string {
	return stringUnixTime
}
