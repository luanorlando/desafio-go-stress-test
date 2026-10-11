package loadtest

import (
	"net/http"
	"sync"
	"time"
)

type Config struct {
	URL         string
	Requests    int
	Concurrency int
}

type Metrics struct {
	TotalDuration time.Duration
	TotalRequests int
	Success200    int64
	StatusDist    map[int]int
}

type Result struct {
	StatusCode int
	Err        error
}

func Run(cfg Config) Metrics {
	startTime := time.Now()

	jobs := make(chan struct{}, cfg.Requests)
	results := make(chan Result, cfg.Requests)

	for i := 0; i < cfg.Requests; i++ {
		jobs <- struct{}{}
	}
	close(jobs)

	var wg sync.WaitGroup

	for w := 0; w < cfg.Concurrency; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{
				Timeout: 10 * time.Second,
			}
			for range jobs {
				resp, err := client.Get(cfg.URL)
				if err != nil {
					results <- Result{StatusCode: 0, Err: err}
					continue
				}
				resp.Body.Close()
				results <- Result{StatusCode: resp.StatusCode, Err: nil}
			}
		}()
	}

	wg.Wait()
	close(results)

	duration := time.Since(startTime)

	metrics := Metrics{
		TotalDuration: duration,
		TotalRequests: cfg.Requests,
		StatusDist:    make(map[int]int),
	}

	var mu sync.Mutex
	for res := range results {
		mu.Lock()
		if res.Err != nil {
			metrics.StatusDist[0]++
		} else {
			metrics.StatusDist[res.StatusCode]++
			if res.StatusCode == http.StatusOK {
				metrics.Success200++
			}
		}
		mu.Unlock()
	}

	return metrics
}
