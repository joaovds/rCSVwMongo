package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandleCollection interface {
  InsertOne(collectionName string, data interface{}) (string, error)
  GetOne(collectionName string, filter bson.D) (map[string]interface{}, error)
  GetMany(collectionName string, filter bson.D) ([]map[string]interface{}, error)
}

func GetMany(collectionName string, filter interface{}) ([]map[string]interface{}, error) {
  collection := GetMongoClient().Database("teste_golang").Collection(collectionName)

  cursor, err := collection.Find(context.TODO(), bson.D{})
  if err != nil {
    return nil, err
  }
  defer cursor.Close(context.TODO())

  var resultsData []map[string]interface{}
  err = cursor.All(context.TODO(), &resultsData)
  if err != nil {
    log.Panicln(err)
  }

  return resultsData, nil
}

func GetOne(collectionName string, filter bson.D) (map[string]interface{}, error) {
  collection := GetMongoClient().Database("teste_golang").Collection(collectionName)

  var resultData map[string]interface{}
  err := collection.FindOne(context.TODO(), filter).Decode(&resultData)
  if err != nil {
    if err == mongo.ErrNoDocuments {
      return nil, nil
    }

    return nil, err
  }

  return resultData, nil
}

