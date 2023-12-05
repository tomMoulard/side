package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func New(host, user, pwd string) (*DB, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/side?ssl=false&authSource=admin", user, pwd, host)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	return &DB{client: client}, nil
}

func (db *DB) Close() error {
	if err := db.client.Disconnect(context.Background()); err != nil {
		return fmt.Errorf("failed to disconnect from mongodb: %w", err)
	}

	return nil
}

func (db *DB) InsertMany(collection string, docs []interface{}) error {
	coll := db.client.Database("side").Collection(collection)
	for _, doc := range docs {
		if _, err := coll.InsertOne(context.Background(), doc); err != nil {
			return fmt.Errorf("failed to insert document: %w", err)
		}
	}

	return nil
}

func (db *DB) Find(collection string, filter bson.D, v []any) error {
	coll := db.client.Database("side").Collection(collection)

	cur, err := coll.Find(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to find documents: %w", err)
	}

	if err := cur.All(context.Background(), &v); err != nil {
		return fmt.Errorf("failed to decode documents: %w", err)
	}

	return nil
}
