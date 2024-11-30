package stresstest

import (
	"encoding/base64"
	"fmt"
	"log"
	"sync"
	"time"
)

type (
	StressTest struct {
		Url         string
		Method      string
		Headers     []string
		BodyEncoded string
		Timeout     time.Duration
		Requests    int
		Concurrency int
	}
)

func (st *StressTest) Run() (*Report, error) {
	var data []byte
	if st.BodyEncoded != "" {
		var err error
		data, err = base64.StdEncoding.DecodeString(st.BodyEncoded)
		if err != nil {
			return nil, fmt.Errorf("invalid body base64: %w", err)
		}
	}

	requestsPerWorker := st.Requests / st.Concurrency
	extraRequests := st.Requests % st.Concurrency

	start := time.Now()
	reports := make(chan Report, st.Concurrency)
	var wg sync.WaitGroup

	for i := 0; i < st.Concurrency; i++ {
		count := requestsPerWorker
		if i < extraRequests {
			count++
		}

		wg.Add(1)
		go func(count int) {
			defer wg.Done()

			r := Report{FailedRequests: make(map[int]int)}
			for i := 0; i < count; i++ {
				code, err := DefaultRequester.MakeRequest(st.Url, st.Method, st.Headers, data, st.Timeout)
				if err != nil {
					log.Printf("Request error: %v", err)
					r.FailedRequests[0]++
					continue
				}
				if code >= 200 && code < 300 {
					r.SuccessfulRequests++
				} else {
					r.FailedRequests[code]++
				}
			}
			reports <- r
		}(count)
	}

	wg.Wait()
	close(reports)

	// Merge all results
	report := Report{FailedRequests: make(map[int]int)}
	for r := range reports {
		report.SuccessfulRequests += r.SuccessfulRequests
		report.TotalRequests += r.SuccessfulRequests
		for k, v := range r.FailedRequests {
			report.FailedRequests[k] += v
			if k > 0 {
				report.TotalRequests += v
			}
		}
	}

	report.TimeSpent = time.Since(start)
	return &report, nil
}
