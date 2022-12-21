package internal

import (
	"context"
	"github.com/andersonribeir0/starfields/pkg"
)

type DBAdapterMock struct {
	SaveFunc func(data interface{}) error
}

func (mock *DBAdapterMock) Save(ctx context.Context, data interface{}) error {
	return mock.SaveFunc(data)
}

type PlanetRepositoryMock struct {
	SaveFunc func(item *Planet) error
}

func (mock *PlanetRepositoryMock) Save(item *Planet) error {
	return mock.SaveFunc(item)
}

type HTTPClientMock struct {
	AsyncGetRequestFunc func(url string, headers map[string][]string, response chan<- *pkg.HTTPResponse)
}

func (mock *HTTPClientMock) AsyncGetRequest(url string,
	headers map[string][]string,
	response chan<- *pkg.HTTPResponse,
) {
	mock.AsyncGetRequestFunc(url,
		headers,
		response)
}
