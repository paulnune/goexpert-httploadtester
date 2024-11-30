package stresstest

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_Request_OK(t *testing.T) {
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer svc.Close()

	code := DefaultRequester.MakeRequest(svc.URL, http.MethodGet, nil, nil, time.Second*5)

	assert.Equal(t, http.StatusOK, code)
}

func Test_Request_BadRequest(t *testing.T) {
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer svc.Close()

	code := DefaultRequester.MakeRequest(svc.URL, http.MethodGet, nil, nil, time.Second*5)

	assert.Equal(t, http.StatusBadRequest, code)
}

func Test_Request_Fail(t *testing.T) {
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	svc.Close()

	code := DefaultRequester.MakeRequest(svc.URL, http.MethodGet, nil, nil, time.Second*5)

	assert.Equal(t, 0, code)
}

func Test_Request_Timeout(t *testing.T) {
	var svc *httptest.Server
	svc = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 1)
	}))
	defer svc.Close()

	start := time.Now()
	code := DefaultRequester.MakeRequest(svc.URL, http.MethodGet, nil, nil, time.Millisecond*500)
	elapsed := time.Since(start)

	assert.Equal(t, 0, code)
	assert.LessOrEqual(t, elapsed, time.Millisecond*600)
}
