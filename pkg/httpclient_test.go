package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestHttpClient(t *testing.T) {
	t.Parallel()

	for scenario, fn := range map[string]func(t *testing.T){
		"invalid_url":         testHTTPClientInvalidURL,
		"internal_server_err": testHTTPClientInternalServerError,
		"success":             testHTTPClientDoSuccessfully,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testHTTPClientInvalidURL(t *testing.T) {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	defer srv.Close()

	httpClient := NewHTTPClient()
	resultChan := make(chan *HTTPResponse)

	httpClient.AsyncGetRequest("...", nil, resultChan)

	result := <-resultChan

	assert.Error(t, result.Err())
}

func testHTTPClientInternalServerError(t *testing.T) {
	t.Helper()

	resultChan := make(chan *HTTPResponse)
	expectedStatusCode := http.StatusInternalServerError
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var js json.RawMessage

		bReq, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(bReq, &js); err != nil {
			w.WriteHeader(expectedStatusCode)
			w.Write([]byte(err.Error()))
		}

		return
	}))

	defer srv.Close()

	httpClient := NewHTTPClient()
	httpClient.AsyncGetRequest(srv.URL, map[string][]string{"foo": {"bar"}}, resultChan)

	result := <-resultChan

	assert.Equal(t, expectedStatusCode, result.StatusCode())
}

func testHTTPClientDoSuccessfully(t *testing.T) {
	t.Helper()

	resultChan := make(chan *HTTPResponse)
	expected := uuid.New().String()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))

	defer srv.Close()

	httpClient := NewHTTPClient()
	httpClient.AsyncGetRequest(srv.URL, map[string][]string{"foo": {"bar"}}, resultChan)

	result := <-resultChan

	assert.Equal(t, http.StatusOK, result.StatusCode())
	assert.Equal(t, expected, string(result.body))
}
