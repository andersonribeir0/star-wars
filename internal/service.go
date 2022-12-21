package internal

import (
	"encoding/json"
	"net/http"

	"github.com/andersonribeir0/starfields/pkg"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	pullPlanetByIDFunc = "pull_planet_by_id_func"
	getFilmDetailsFunc = "get_film_details_func"
	swURL              = "https://swapi.dev/api"
)

type PullPlanetResponse struct {
	Name    string   `json:"name"`
	Climate string   `json:"climate"`
	Terrain string   `json:"terrain"`
	Films   []string `json:"films"`
}

type PlanetServiceI interface {
	PullPlanetByID(id string) error
}

type PlanetService struct {
	log              *zap.Logger
	httpClient       pkg.HTTPClientI
	planetRepository PlanetRepositoryI
}

func NewPlanetService(log *zap.Logger,
	httpClient pkg.HTTPClientI,
	planetRepository PlanetRepositoryI,
) *PlanetService {
	return &PlanetService{
		log:              log,
		httpClient:       httpClient,
		planetRepository: planetRepository,
	}
}

func (service *PlanetService) PullPlanetByID(id string) error {
	var (
		respBody       *PullPlanetResponse
		films          []*Film
		httpPlanetChan = make(chan *pkg.HTTPResponse)
		filmsResp      = make(chan *pkg.HTTPResponse)
	)

	if id == "" {
		return ErrPlanetIDEmpty
	}

	go service.httpClient.AsyncGetRequest(
		swURL+"/planets/"+id,
		nil,
		httpPlanetChan,
	)

	resp := <-httpPlanetChan

	if resp.Err() != nil {
		return errors.Wrap(resp.Err(), pullPlanetByIDFunc)
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.Wrap(ErrHTTPStatusCode, pullPlanetByIDFunc)
	}

	err := json.Unmarshal(resp.Body(), &respBody)
	if err != nil {
		return errors.Wrap(err, pullPlanetByIDFunc)
	}

	for _, v := range respBody.Films {
		go service.httpClient.AsyncGetRequest(
			v,
			nil,
			filmsResp,
		)
	}

	for j := 0; j < len(respBody.Films); j++ {
		item, err := service.getFilmDetails(<-filmsResp)
		if err != nil {
			return errors.Wrap(err, pullPlanetByIDFunc)
		}

		films = append(films, item)
	}

	err = service.planetRepository.Save(&Planet{
		ExternalID: id,
		Name:       respBody.Name,
		Climate:    respBody.Climate,
		Terrain:    respBody.Terrain,
		Films:      films,
	})

	return errors.Wrap(err, pullPlanetByIDFunc)
}

func (service *PlanetService) getFilmDetails(resp *pkg.HTTPResponse) (body *Film, err error) {
	if resp.Err() != nil {
		return nil, errors.Wrap(resp.Err(), getFilmDetailsFunc)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.Wrap(ErrHTTPStatusCode, getFilmDetailsFunc)
	}

	err = json.Unmarshal(resp.Body(), &body)

	return body, errors.Wrap(err, getFilmDetailsFunc)
}
