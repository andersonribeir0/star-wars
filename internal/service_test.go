package internal

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/andersonribeir0/starfields/pkg"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestPlanetService(t *testing.T) {
	t.Parallel()

	for scenario, fn := range map[string]func(t *testing.T){
		"pull_planet_id_empty":      testPlanetServicePullPlanetIDEmpty,
		"pull_planet_get_films_err": testPlanetServicePullPlanetGetFilmsErr,
		"pull_planet_empty_body":    testPlanetServicePullPlanetEmptyBody,
		"pull_planet_http_err":      testPlanetServicePullPlanetHTTPErr,
		"pull_planet_successfully":  testPlanetServicePullPlanetSuccessfully,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testPlanetServicePullPlanetIDEmpty(t *testing.T) {
	t.Helper()

	service := NewPlanetService(nil,
		&HTTPClientMock{AsyncGetRequestFunc: setupHTTPResponseGetFilmsHTTPErr},
		&PlanetRepositoryMock{SaveFunc: setupSave},
	)

	assert.Equal(t, ErrPlanetIDEmpty, errors.Cause(service.PullPlanetByID("")))
}

func testPlanetServicePullPlanetGetFilmsErr(t *testing.T) {
	t.Helper()

	log, err := zap.NewDevelopment()
	assert.NoError(t, err)

	service := NewPlanetService(log,
		&HTTPClientMock{AsyncGetRequestFunc: setupHTTPResponseGetFilmsHTTPErr},
		&PlanetRepositoryMock{SaveFunc: setupSave},
	)

	err = service.PullPlanetByID(uuid.NewString())
	assert.Equal(t, ErrHTTPStatusCode, errors.Cause(err))

	service = NewPlanetService(log,
		&HTTPClientMock{AsyncGetRequestFunc: setupHTTPResponseGetFilmsErr},
		&PlanetRepositoryMock{SaveFunc: setupSave},
	)

	err = service.PullPlanetByID(uuid.NewString())
	assert.Error(t, err)
}

func testPlanetServicePullPlanetEmptyBody(t *testing.T) {
	t.Helper()

	log, err := zap.NewDevelopment()
	assert.NoError(t, err)

	service := NewPlanetService(log,
		&HTTPClientMock{AsyncGetRequestFunc: setupHTTPResponseEmptyBody},
		&PlanetRepositoryMock{SaveFunc: setupSave},
	)

	err = service.PullPlanetByID(uuid.NewString())
	assert.Error(t, err)
}

func testPlanetServicePullPlanetHTTPErr(t *testing.T) {
	t.Helper()

	log, err := zap.NewDevelopment()
	assert.NoError(t, err)

	service := NewPlanetService(log,
		&HTTPClientMock{AsyncGetRequestFunc: setupHTTPResponseHTTPStatusCodeErr},
		&PlanetRepositoryMock{SaveFunc: setupSave},
	)

	err = service.PullPlanetByID(uuid.NewString())
	assert.Equal(t, ErrHTTPStatusCode, errors.Cause(err))

	service = NewPlanetService(log,
		&HTTPClientMock{AsyncGetRequestFunc: setupHTTPResponseHTTPErr},
		&PlanetRepositoryMock{SaveFunc: setupSave},
	)

	err = service.PullPlanetByID(uuid.NewString())
	assert.Error(t, err)
}

func testPlanetServicePullPlanetSuccessfully(t *testing.T) {
	t.Helper()

	log, err := zap.NewDevelopment()
	assert.NoError(t, err)

	service := NewPlanetService(log,
		&HTTPClientMock{AsyncGetRequestFunc: setupHTTPResponseSuccess},
		&PlanetRepositoryMock{SaveFunc: setupSave},
	)

	err = service.PullPlanetByID(uuid.NewString())
	assert.NoError(t, err)
}

func setupHTTPResponseGetFilmsHTTPErr(url string, headers map[string][]string, response chan<- *pkg.HTTPResponse) {
	var bResp []byte

	if strings.Contains(url, "planet") {
		bResp, _ = json.Marshal(PullPlanetResponse{
			Name:    uuid.NewString(),
			Climate: uuid.NewString(),
			Terrain: uuid.NewString(),
			Films:   []string{"https://localhost/2"},
		})
		response <- pkg.NewHTTPResponse(http.StatusOK, nil, bResp, nil)
	} else {
		bResp, _ = json.Marshal(Film{
			Title:    uuid.NewString(),
			Director: uuid.NewString(),
			Created:  time.Now().UTC(),
		})
		response <- pkg.NewHTTPResponse(http.StatusInternalServerError, nil, bResp, nil)
	}
}

func setupHTTPResponseGetFilmsErr(url string, headers map[string][]string, response chan<- *pkg.HTTPResponse) {
	var bResp []byte

	if strings.Contains(url, "planet") {
		bResp, _ = json.Marshal(PullPlanetResponse{
			Name:    uuid.NewString(),
			Climate: uuid.NewString(),
			Terrain: uuid.NewString(),
			Films:   []string{"https://localhost/1"},
		})
		response <- pkg.NewHTTPResponse(http.StatusOK, nil, bResp, nil)
	} else {
		bResp, _ = json.Marshal(Film{
			Title:    uuid.NewString(),
			Director: uuid.NewString(),
			Created:  time.Now().UTC(),
		})
		response <- pkg.NewHTTPResponse(http.StatusOK, errors.New("err"), bResp, nil)
	}
}

func setupHTTPResponseEmptyBody(url string, headers map[string][]string, response chan<- *pkg.HTTPResponse) {
	response <- pkg.NewHTTPResponse(http.StatusOK, nil, nil, nil)
}

func setupHTTPResponseHTTPErr(url string, headers map[string][]string, response chan<- *pkg.HTTPResponse) {
	response <- pkg.NewHTTPResponse(http.StatusOK, errors.New("err"), nil, nil)
}

func setupHTTPResponseHTTPStatusCodeErr(url string, headers map[string][]string, response chan<- *pkg.HTTPResponse) {
	response <- pkg.NewHTTPResponse(http.StatusInternalServerError, nil, nil, nil)
}

func setupHTTPResponseSuccess(url string, headers map[string][]string, response chan<- *pkg.HTTPResponse) {
	var bResp []byte

	if strings.Contains(url, "planet") {
		bResp, _ = json.Marshal(PullPlanetResponse{
			Name:    uuid.NewString(),
			Climate: uuid.NewString(),
			Terrain: uuid.NewString(),
			Films:   []string{"https://localhost/1"},
		})
	} else {
		bResp, _ = json.Marshal(Film{
			Title:    uuid.NewString(),
			Director: uuid.NewString(),
			Created:  time.Now().UTC(),
		})
	}

	response <- pkg.NewHTTPResponse(http.StatusOK, nil, bResp, nil)
}

func setupSave(item *Planet) error {
	return nil
}
