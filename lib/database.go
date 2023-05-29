package lib

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseManagement interface {
	NewDatabase(db_name string)
	ConnectDB(collection_name string)
	FindOneQuery(filter bson.D) bson.M
	DeleteOneQuery(filter bson.D) bool
}

type Database struct {
	DBUri                   string
	DBClient                *mongo.Client
	DBCurrentCollection     *mongo.Collection
	DBName                  string
	DBCurrentCollectionName string
}

// CREATE NEW DB
func NewDatabase(db_name string) Database {
	if db_name == "" {
		db_name = "panel"
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return Database{
		DBUri:                   os.Getenv("DB_URI"),
		DBClient:                &mongo.Client{},
		DBCurrentCollection:     &mongo.Collection{},
		DBName:                  db_name,
		DBCurrentCollectionName: "",
	}
}

// DB CONNECTION
func (db *Database) ConnectDB(collection_name string) *mongo.Collection {

	db.DBCurrentCollectionName = collection_name

	if db.DBUri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBUri))
	if err != nil {
		panic(err)
	}

	db.DBClient = client

	collection := client.Database(db.DBName).Collection(db.DBCurrentCollectionName)
	db.DBCurrentCollection = collection

	return collection
}

// DB OPERATIONS

func (db *Database) FindOneUser(document bson.D) (User, error) {
	var result User
	err := db.DBCurrentCollection.FindOne(context.TODO(), document).Decode(&result)

	if err == mongo.ErrNoDocuments {
		Logger(errors.New(fmt.Sprintf("ERR: Document not found in the collection: %s", db.DBCurrentCollection)))
		return User{
			UserID:   primitive.NewObjectID(),
			Type:     false,
			Username: "NOT_FOUND",
			Password: "NOT_FOUND",
		}, errors.New("ERROR DOCUMENT NOT FOUNDED!")
	}
	Logger(err)

	return result, nil
}

func (db *Database) FindOneDocument(document bson.D) (Document, error) {
	var result Document

	err := db.DBCurrentCollection.FindOne(context.TODO(), document).Decode(&result)

	if err == mongo.ErrNoDocuments {
		Logger(errors.New(fmt.Sprintf("ERR: Document not found in the collection: %s", db.DBCurrentCollection)))
		return Document{
			ID:         primitive.NewObjectID(),
			Content:    "NOT_FOUND",
			Author:     "NOT_FOUND",
			AccessList: []primitive.ObjectID{},
		}, errors.New("ERROR DOCUMENT NOT FOUNDED!")
	}
	Logger(err)

	return result, nil
}

func (db *Database) DeleteOneQuery(filter bson.D) bool {
	_, err := db.DBCurrentCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		Logger(errors.New("WE CAN'T DELETE DOCUMENT FROM DB FOR NOW!"))
		return false
	}
	return true
}
