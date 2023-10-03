package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/joaovds/rCSVwMongo/internal/database"
	"github.com/joaovds/rCSVwMongo/internal/entities"
	"github.com/joaovds/rCSVwMongo/pkg/handleCSV"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func loadEnv() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file:", err)
  }
}

func main() {
  loadEnv()

  csvData, err := handleCSV.ReadCSV("./assets/csvTest.csv")
  if err != nil {
    log.Fatalln("Error reading CSV file:", err)
  }

  data := handleCSV.ConvertToStruct(csvData)

  jsonEncode(data)

  resultsData, err := database.GetMany("tests_rCSVwMONGO", bson.D{})
  if err != nil {
    log.Panicln("Error getting data from database:", err)
  }

  resultData, err := database.GetOne("tests_rCSVwMONGO", bson.D{
    primitive.E{Key: "csvID", Value: "1"},
    primitive.E{Key: "email", Value: "teste@teste.com"},
  })
  if err != nil {
    log.Panicln("Error getting data from database:", err)
  }

  resultCreate, err := database.InsertOne("tests_rCSVwMONGO", entities.User{
    CsvId: "1",
    Name: "Teste",
    Email: "teste@teste2.com",
    Age: 20,
    Date: time.Now(),
  })
  if err != nil {
    log.Panicln("Error inserting data from database:", err)
  }

  fmt.Println(resultsData)
  fmt.Println(resultData)
  fmt.Println(resultCreate)
}

func jsonEncode(data []entities.User) {
  json, err := json.MarshalIndent(data, "", "  ")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(json))
}

