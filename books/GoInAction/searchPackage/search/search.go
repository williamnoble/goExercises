package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel for the results
	results := make(chan *Result)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			//	matcher = matcher["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			//		Match(matcher, feed, searchTerm, results)
			waitGroup.Done()

		}(matcher, feed)

		go func() {
			waitGroup.Wait()
			close(results)
		}()

		//	Display(results)
	}

}
