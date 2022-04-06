package db

import (
	"context"

	"backend-food/internal/pkg/domain/domain_model/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	DBName string
}

func NewDB() (Database, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://hiendv021000:leloi123456@cluster0.lqike.mongodb.net/myFirstDatabase?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	return Database{Client: client,
		DBName: "food",
	}, err
}

func (db *Database) First(condition interface{}, value entity.BaseEntity) error {
	collection := db.Client.Database(db.DBName).Collection(value.CollectionName())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, condition).Decode(value)
	return err
}
func (db *Database) Find(condition interface{}, model entity.BaseEntity) ([]bson.M, error) {
	collection := db.Client.Database(db.DBName).Collection(model.CollectionName())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, condition)
	if err != nil {
		return nil, err
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		return nil, err
	}

	return episodes, nil
}
func (db *Database) Create(value entity.BaseEntity) error {
	collection := db.Client.Database(db.DBName).Collection(value.CollectionName())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, value)
	if err != nil {
		return err
	}
	value.SetID(res.InsertedID.(primitive.ObjectID))
	return err
}
func (db *Database) Delete(value entity.BaseEntity) error {
	collection := db.Client.Database(db.DBName).Collection(value.CollectionName())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, value)
	return err
}
func (db *Database) Update(filter entity.BaseEntity, update entity.BaseEntity) error {
	collection := db.Client.Database(db.DBName).Collection(filter.CollectionName())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, filter, bson.D{
		{"$set", update},
	})
	return err
}
