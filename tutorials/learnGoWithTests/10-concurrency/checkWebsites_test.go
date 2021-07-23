package concurrency

import (
	"reflect"
	"testing"
	"time"
)

var websites = []string{
	"http://www.google.com",
	"http://www.bing.com",
	"waat://furhuterewe.ged",
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhuterewe.ged" {
		return false
	}
	return true
}

func slowStudWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true

}
func TestCheckWebsites(t *testing.T) {
	want := map[string]bool{
		"http://www.google.com":  true,
		"http://www.bing.com":    true,
		"waat://furhuterewe.ged": false,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStudWebsiteChecker, urls)
	}
}
