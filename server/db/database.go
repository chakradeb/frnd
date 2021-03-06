package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/chakradeb/frnd-server/models"
)

type IDBClient interface {
	CreateUser (string, string) error
	GetUser(string) (*models.User, error)
	GetProfile(string) (*models.Profile, error)
}

type DB struct {
	database *mongo.Database
	ctx context.Context
}

func New(dbHost string, dbPort int, dbName string) (*DB, error) {
	connectionString := fmt.Sprintf("mongodb://%s:%d", dbHost, dbPort)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("db: error creating client on %s", connectionString)
	}

	ctx := context.TODO()
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	if err != nil {
		return nil, fmt.Errorf("db: unable to connect on host %s: %s", connectionString, err)
	}

	return &DB{
		database: client.Database(dbName),
		ctx: ctx,
	}, nil
}

func (d *DB) CreateUser(username string, password string) error {
	_, err := d.database.Collection("users").InsertOne(d.ctx, bson.D{
		{Key: "username", Value: username},
		{Key: "password", Value: password},
	})
	_, err = d.database.Collection("profiles").InsertOne(d.ctx, bson.D{
		{Key: "username", Value: username},
		{Key: "followers", Value: 0},
	})
	return err
}

func (d *DB) GetUser(username string) (*models.User, error) {
	user := &models.User{}
	filter := bson.M{"username": username}

	err := d.database.Collection("users").FindOne(d.ctx, filter).Decode(user)
	if err != nil {
		return nil, fmt.Errorf("db: unable to get user %s: %s", username, err)
	}
	return user, nil
}

func (d *DB) GetProfile(username string) (*models.Profile, error) {
	profile := &models.Profile{}
	filter := bson.M{"username": username}

	err := d.database.Collection("profiles").FindOne(d.ctx, filter).Decode(profile)
	if err != nil {
		return nil, fmt.Errorf("db: unable to get profile %s: %s", username, err)
	}
	return profile, nil
}
