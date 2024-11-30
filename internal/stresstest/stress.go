package stresstest

import (
	"encoding/base64"
	"fmt"
	"log"
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

	executionsPerFork := st.Requests / st.Concurrency
	remainingRequests := st.Requests % st.Concurrency

	start := time.Now()
	reports := make(chan Report)
	for i := 0; i < st.Concurrency; i++ {
		count := executionsPerFork
		if remainingRequests > 0 {
			remainingRequests--
			count++
		}

		log.Printf("Requests per batch: %d, Concurrency: %d", count, st.Concurrency)

		go func(count int) {
			r := Report{
				FailedRequests: make(map[int]int),
			}
			for i := 0; i < count; i++ {
				code := DefaultRequester.MakeRequest(st.Url, st.Method, st.Headers, data, st.Timeout)
				if code >= 200 && code < 300 {
					r.SuccessfulRequests++
				} else {
					r.FailedRequests[code]++
				}
			}
			reports <- r
		}(count)
	}

	// merge all results
	report := Report{
		FailedRequests: make(map[int]int),
	}
	for i := 0; i < st.Concurrency; i++ {
		r := <-reports
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
