package internal

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestPlanetRepository(t *testing.T) {
	t.Parallel()

	for scenario, fn := range map[string]func(t *testing.T){
		"save_planet_failed":       testSavePlanetFailed,
		"save_planet_successfully": testSavePlanetSuccessfully,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testSavePlanetFailed(t *testing.T) {
	t.Helper()

	log, err := zap.NewDevelopment()
	assert.NoError(t, err)

	repo := NewPlanetRepository(log, &DBAdapterMock{SaveFunc: func(data interface{}) error {
		return errors.New("err")
	}})

	assert.Error(t, repo.Save(&Planet{}))
}

func testSavePlanetSuccessfully(t *testing.T) {
	t.Helper()

	log, err := zap.NewDevelopment()
	assert.NoError(t, err)

	repo := NewPlanetRepository(log, &DBAdapterMock{SaveFunc: func(data interface{}) error {
		return nil
	}})

	assert.NoError(t, repo.Save(&Planet{
		ExternalID: uuid.NewString(),
		Name:       uuid.NewString(),
		Climate:    uuid.NewString(),
		Terrain:    uuid.NewString(),
		Films: []*Film{{
			Title:    uuid.NewString(),
			Director: uuid.NewString(),
			Created:  time.Now().UTC(),
		}},
	}))
}
