package internal

import (
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Planet struct {
	Name    string  `json:"name" bson:"name"`
	Climate string  `json:"climate" bson:"climate"`
	Terrain string  `json:"terrain" bson:"terrain"`
	Films   []*Film `json:"films" bson:"films"`
}

type Film struct {
	Title    string    `json:"title" bson:"title"`
	Director string    `json:"director" bson:"director"`
	Created  time.Time `json:"created" bson:"created"`
}

type PlanetRepositoryI interface {
	Save(item *Planet) error
}

type PlanetRepository struct {
	log *zap.Logger
}

func NewPlanetRepository(log *zap.Logger) *PlanetRepository {
	return &PlanetRepository{log: log}
}

func (pr *PlanetRepository) Save(item *Planet) error {
	bResp, _ := json.Marshal(item)
	pr.log.Info(fmt.Sprintf("SAVING %s", string(bResp)))

	return nil
}
