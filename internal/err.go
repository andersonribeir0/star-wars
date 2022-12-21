package internal

import "github.com/pkg/errors"

var (
	ErrPlanetIDEmpty  = errors.New("planet id must be provided")
	ErrHTTPStatusCode = errors.New("http status code is not 200")
)
