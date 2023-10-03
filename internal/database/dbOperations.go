package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/joaovds/rCSVwMongo/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func CheckIfExist(collectionName string, filter bson.D) (bool, error) {
  document, err := GetOne(collectionName, filter)
  if err != nil {
    return false, err
  }
  
  if document == nil {
    return false, nil
  }

  return true, nil
}

func InsertOne(collectionName string, data interface{}) (string, error) {
  dataTypeAsserted, ok := data.(entities.User)
  if !ok {
    return "", errors.New("Erro ao converter o tipo de dado")
  }

  exists, err := CheckIfExist(
    collectionName,
    bson.D{
      primitive.E{ Key: "csvID", Value: dataTypeAsserted.CsvId },
      primitive.E{ Key: "email", Value: dataTypeAsserted.Email },
    },
    )
  if err != nil {
    log.Panicln(err)
  }

  if exists {
    str := fmt.Sprintf("Usuário com o ID: %s e Email: %s já existe", dataTypeAsserted.CsvId, dataTypeAsserted.Email)
    return str, nil
  } else {
    collection := GetMongoClient().Database("teste_golang").Collection(collectionName)
    _, err := collection.InsertOne(context.TODO(), data)
    if err != nil {
      return "", err
    }

    successMessage := fmt.Sprintf("Inserido com Sucesso: %s", dataTypeAsserted.CsvId)
    return successMessage, nil
  }
}

