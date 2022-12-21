package internal

import (
	"context"
	"github.com/andersonribeir0/starfields/pkg"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"time"
)

const (
	dbTimeout        = time.Second * 15
	newDBAdapterFunc = "new_db_adapter_func"
)

type DBAdapterI interface {
	Save(ctx context.Context, data interface{}) error
}

type DBAdapter struct {
	client   *mongo.Client
	database *mongo.Database
	log      *zap.Logger
}

func NewDBAdapter(connectionS string, log *zap.Logger) (*DBAdapter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionS))
	if err != nil {
		return nil, errors.Wrap(err, newDBAdapterFunc)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, errors.Wrap(err, newDBAdapterFunc)
	}

	log.Info("db connected")

	return &DBAdapter{
		client:   client,
		database: client.Database("star-wars"),
		log:      log,
	}, nil
}

func (adapter *DBAdapter) Save(ctx context.Context, data interface{}) (err error) {
	collection := adapter.database.Collection(pkg.TypeName(data))

	_, err = collection.InsertOne(ctx, data)

	return errors.Wrap(err, "repository_save")
}
