package internal

import (
	"context"

	"github.com/andersonribeir0/starfields/pkg"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Components are a like service, but it doesn't include business case
// Or domains, but likely used by multiple domains
type components struct {
	Log *zap.Logger
	// Include your new components bellow
	HttpClient       *pkg.HTTPClient
	PlanetRepository PlanetRepositoryI
}

// Services hold the business case, and make the bridge between
// Controllers and Domains
type Services struct {
	// Include your new services bellow
	PlanetService PlanetServiceI
}

type Dependency struct {
	Components components
	Services   Services
}

func NewContainer(ctx context.Context) (*Dependency, error) {
	cmp, err := setupComponents(ctx)
	if err != nil {
		return nil, err
	}

	srv := setupServices(ctx, cmp)

	dep := Dependency{
		Components: *cmp,
		Services:   *srv,
	}

	return &dep, err
}

func setupServices(ctx context.Context, cmp *components) *Services {
	return &Services{PlanetService: NewPlanetService(cmp.Log, cmp.HttpClient, cmp.PlanetRepository)}
}

func setupComponents(ctx context.Context) (*components, error) {
	log, err := setupLog()
	if err != nil {
		return nil, err
	}

	planetRepository := NewPlanetRepository(log)

	return &components{
		Log:              log,
		HttpClient:       pkg.NewHTTPClient(),
		PlanetRepository: planetRepository,
	}, nil
}

func setupLog() (*zap.Logger, error) {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	loggerConfig.DisableStacktrace = false
	loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	logger, err := loggerConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, errors.Wrap(err, "error on building zap logger")
	}

	return logger, nil
}
