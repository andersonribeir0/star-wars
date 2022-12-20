package pkg

import (
	"bytes"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	asyncRequestFunc = "async_request_func"
)

type HTTPResponse struct {
	statusCode int
	err        error
	body       []byte
	headers    map[string][]string
}

func (h *HTTPResponse) StatusCode() int {
	return h.statusCode
}

func (h *HTTPResponse) Err() error {
	return h.err
}

func (h *HTTPResponse) Body() []byte {
	return h.body
}

func (h *HTTPResponse) Headers() map[string][]string {
	return h.headers
}

type HTTPClient struct{}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{}
}

func (h *HTTPClient) AsyncRequest(method, url string, headers map[string][]string, body []byte) (resp <-chan *HTTPResponse) {
	var (
		err      error
		req      *http.Request
		response = make(chan *HTTPResponse)
	)

	go func() {
		defer close(response)

		if body == nil {
			req, err = http.NewRequest(method, url, http.NoBody)
		} else {
			bodyReader := bytes.NewReader(body)
			req, err = http.NewRequest(method, url, bodyReader)
		}

		for headerKey, headerValues := range headers {
			for _, headerValue := range headerValues {
				req.Header.Add(headerKey, headerValue)
			}
		}

		if err != nil {
			response <- &HTTPResponse{
				statusCode: http.StatusInternalServerError,
				err:        errors.Wrap(err, asyncRequestFunc),
				body:       nil,
				headers:    nil,
			}

			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			response <- &HTTPResponse{
				statusCode: http.StatusInternalServerError,
				err:        errors.Wrap(err, asyncRequestFunc),
				body:       nil,
				headers:    nil,
			}

			return
		}

		defer res.Body.Close()

		bodyResponse, err := io.ReadAll(res.Body)
		if err != nil {
			response <- &HTTPResponse{
				statusCode: http.StatusInternalServerError,
				err:        errors.Wrap(err, asyncRequestFunc),
				body:       nil,
				headers:    nil,
			}

			return
		}

		response <- &HTTPResponse{
			statusCode: res.StatusCode,
			err:        nil,
			body:       bodyResponse,
			headers:    res.Header,
		}
	}()

	return response
}
