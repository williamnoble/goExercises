package _1_select_partone

import (
	"net/http"
	"time"
)

func OriginalRacer(a, b string) (winner string) {
	aDuration, bDuration := OriginalMeasureresponsetime(a), OriginalMeasureresponsetime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func OriginalMeasureresponsetime(url string) time.Duration {
	start := time.Now()
	//goland:noinspection GoUnhandledErrorResult
	http.Get(url)
	return time.Since(start)
}
