package database

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
  once sync.Once
  mongoDb *mongo.Client
)

func GetMongoClient() *mongo.Client {
  once.Do(func() {
    databaseURI := os.Getenv("DATABASE_URL")
    if databaseURI == "" {
      log.Fatal("DATABASE_URL is empty")
    }

    mongoOptions := options.Client().ApplyURI(databaseURI)
    db, err := mongo.Connect(context.TODO(), mongoOptions)
    if err != nil {
      panic(err)
    }

    mongoDb = db
  })

  return mongoDb
}

