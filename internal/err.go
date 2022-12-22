package internal

import "github.com/pkg/errors"

const errUnknown = "UNKNOWN"

var (
	ErrPlanetIDEmpty       = errors.New("planet id must be provided")
	ErrHTTPStatusCode      = errors.New("http status code is not 200")
	ErrPlanetAlreadyExists = errors.New("planet already exists")
)

func ErrorCode(err error) string {
	switch err != nil {
	case errors.Is(err, ErrPlanetAlreadyExists):
		return "PLANET_ALREADY_EXISTS"
	case errors.Is(err, ErrHTTPStatusCode):
		return "INTEGRATION_ERR"
	case errors.Is(err, ErrPlanetIDEmpty):
		return "PLANET_ID_EMPTY"
	default:
		return errUnknown
	}
}
