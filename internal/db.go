package internal

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/andersonribeir0/starfields/pkg"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

const (
	dbTimeout               = time.Second * 15
	newDBAdapterFunc        = "new_db_adapter_func"
	collectionAlreadyExists = 48
	documentAlreadyExists   = 11000
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

	dbName := "star-wars"
	docName := pkg.TypeName(&Planet{})

	err = client.Database(dbName).CreateCollection(context.TODO(),
		docName)
	if err != nil {
		mongoErr := err.(mongo.CommandError)

		if mongoErr.Code == collectionAlreadyExists {
			return &DBAdapter{
				client:   client,
				database: client.Database(dbName),
				log:      log,
			}, nil
		}

		return nil, errors.Wrap(err, newDBAdapterFunc)
	}

	unique := true
	mod := mongo.IndexModel{
		Keys: bson.M{
			"external_id": 1,
		},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	}

	_, err = client.Database(dbName).Collection(docName).Indexes().CreateOne(context.TODO(), mod)
	if err != nil {
		return nil, errors.Wrap(err, newDBAdapterFunc)
	}

	return &DBAdapter{
		client:   client,
		database: client.Database(dbName),
		log:      log,
	}, nil
}

func (adapter *DBAdapter) Save(ctx context.Context, data interface{}) (err error) {
	collection := adapter.database.Collection(pkg.TypeName(data))

	_, err = collection.InsertOne(ctx, data)

	return errors.Wrap(err, "repository_save")
}
