package checker

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type CheckResponse struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

func CheckURLs(urls []string) []CheckResponse {
	results := make(chan CheckResponse, len(urls)) //buffered channel
	timeout := 2 * time.Second
	var finalResults []CheckResponse

	for _, url := range urls {
		go func(url string) {
			ctx, cancel := context.WithTimeout(context.TODO(), timeout)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				results <- CheckResponse{URL: url, Status: "Error"}
				return
			}

			start := time.Now()
			response, err := http.DefaultClient.Do(req)
			elapsed := time.Since(start)

			if err != nil {
				if ctx.Err() == context.DeadlineExceeded {
					results <- CheckResponse{URL: url, Status: "Time exceed"}
				} else {
					results <- CheckResponse{URL: url, Status: "error"}
				}
				return
			}
			defer response.Body.Close()

			results <- CheckResponse{
				URL:    url,
				Status: fmt.Sprintf("%d %s (%dms)", response.StatusCode, http.StatusText(response.StatusCode), elapsed.Milliseconds()),
			}
		}(url)

	}
	for i := 0; i < len(urls); i++ {
		finalResults = append(finalResults, <-results)
	}
	return finalResults
}
