package internal

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"go.uber.org/zap"
)

type Planet struct {
	ExternalID string  `json:"external_id" bson:"external_id"`
	Name       string  `json:"name" bson:"name"`
	Climate    string  `json:"climate" bson:"climate"`
	Terrain    string  `json:"terrain" bson:"terrain"`
	Films      []*Film `json:"films" bson:"films"`
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
	db  DBAdapterI
}

func NewPlanetRepository(log *zap.Logger, db DBAdapterI) *PlanetRepository {
	return &PlanetRepository{log: log, db: db}
}

func (pr *PlanetRepository) Save(item *Planet) error {
	return errors.Wrap(pr.db.Save(context.TODO(), item), "repository_save")
}
