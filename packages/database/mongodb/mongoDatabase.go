package mongodb

import (
	"context"
	"errors"
	"go-full-cqrs-architecture/packages/database"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDatabase struct {
	mongoClient *mongo.Client
	connString  string
	database    string
}

func NewMongoDatabase(connString string, database string) (database.Database, error) {
	var once sync.Once
	var instanceError error
	var clientInstance *mongo.Client
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(connString)

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			instanceError = err
		}
		err = client.Ping(context.TODO(), readpref.Primary())
		if err != nil {
			instanceError = err
		}
		clientInstance = client
	})

	return &mongoDatabase{
		mongoClient: clientInstance,
		connString:  connString,
		database:    database,
	}, instanceError
}

func (md mongoDatabase) Insert(data database.Data) error {
	client := md.mongoClient
	collection := client.Database(md.database).Collection(data.GetCollection())
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func (md mongoDatabase) InsertMany(data []database.Data) error {

	if len(data) == 0 {
		return errors.New("slice is empty")
	}
	collectionName := data[0].GetCollection()

	client := md.mongoClient
	collection := client.Database(md.database).Collection(collectionName)

	insertableList := make([]interface{}, len(data))
	for i, v := range data {
		insertableList[i] = v
	}

	_, err := collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return err
	}
	return nil
}

func (md mongoDatabase) Update(data database.Data) error {
	client := md.mongoClient
	collection := client.Database(md.database).Collection(data.GetCollection())
	_, err := collection.UpdateByID(context.TODO(), data.GetCollection(), data)
	if err != nil {
		return err
	}
	return nil
}

func (md mongoDatabase) Get(data database.Data) (interface{}, error) {
	client := md.mongoClient
	collection := client.Database(md.database).Collection(data.GetCollection())
	filter := bson.D{primitive.E{Key: "_id", Value: data.GetID()}}

	var result interface{}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
