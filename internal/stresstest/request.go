package stresstest

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type (
	Requester interface {
		MakeRequest(url, method string, headers []string, data []byte, timeout time.Duration) (int, error)
	}
	defaultRequester struct{}
)

var DefaultRequester Requester = &defaultRequester{}

func (_ *defaultRequester) MakeRequest(url, method string, headers []string, data []byte, timeout time.Duration) (int, error) {
	// Configuração para ignorar o certificado TLS
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: timeout,
	}

	var body io.Reader
	if data != nil {
		body = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return 0, fmt.Errorf("error creating request: %w", err)
	}

	if headers != nil {
		for _, h := range headers {
			parts := strings.Split(h, ":")
			if len(parts) != 2 {
				return 0, fmt.Errorf("invalid header format: %s", h)
			}
			req.Header.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error performing request: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
