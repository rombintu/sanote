package store

import (
	"context"
	"fmt"
	"time"

	"github.com/rombintu/sanote/tools"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseSanote string = "sanote"
)

const (
	notesColl string = "notes"
	usersColl string = "users"
)

const (
	noteType    string = "note"
	dirType     string = "directory"
	storageType string = "storage"
)

type Store struct {
	Driver   *mongo.Client
	Database *mongo.Database
	Options  *options.ClientOptions
}

func NewStore() *Store {
	return &Store{
		Options: getOptions(),
	}
}

func getOptions() *options.ClientOptions {
	return options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s",
		tools.GetEnvOrDefault("MONGO_USER", "mongo"),
		tools.GetEnvOrDefault("MONGO_PASS", "mongo"),
		tools.GetEnvOrDefault("MONGO_HOST", "localhost"),
		tools.GetEnvOrDefault("MONGO_PORT", "27017"),
	),
	)
}

// mongodb://<username>:<password>@<host>:<port>
func (s *Store) Open() (context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, s.Options)
	if err != nil {
		return nil, err
	}
	s.Driver = client
	s.Database = s.Driver.Database(databaseSanote)
	return ctx, nil
}

func (s *Store) Close(ctx context.Context) error {
	return s.Driver.Disconnect(ctx)
}
