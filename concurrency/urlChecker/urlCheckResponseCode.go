package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{"https://rogerwelin.github.io/",
		"https://golang.org/",
		"https://news.ycombinator.com/",
		"https://www.google.se/shouldbe404",
		"https://www.cpan.org/"}

	respStatus := make(map[string]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Errorf("Url was Error: %v", err)
			}
			mu.Lock()
			respStatus[url] = resp.StatusCode
			mu.Unlock()
		}(url)

	}
	wg.Wait()
	for key, code := range respStatus {
		fmt.Printf("%s -> %d\n", key, code)
	}
}
