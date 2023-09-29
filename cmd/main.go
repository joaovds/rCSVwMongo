package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CsvData struct {
  Id    string    `json:"id"`
  Name  string    `json:"name"`
  Email string    `json:"email"`
  Age   int       `json:"age"`
  Date  time.Time `json:"date"`
}

func loadEnv() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file:", err)
  }
}

func main() {
  loadEnv()

  file, err := os.Open("./assets/csvTest.csv")
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  reader := csv.NewReader(file)
  reader.Comma = 59
  csvData, err := reader.ReadAll()
  if err != nil {
    log.Fatal(err)
  }

  data := convertToStruct(csvData)

  jsonEncode(data)

  databaseURI := os.Getenv("DATABASE_URL")
  if databaseURI == "" {
    log.Fatal("DATABASE_URL is empty")
  }

  db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseURI))
  if err != nil {
    panic(err)
  }

  defer func() {
    if err := db.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()

  tests_rCSVwMONGOColl := db.Database("teste_golang").Collection("tests_rCSVwMONGO")
  
  var resultsData []bson.M
  cursor, err := tests_rCSVwMONGOColl.Find(context.TODO(), bson.D{})
  if err != nil {
    log.Fatal(err)
  }

  defer cursor.Close(context.TODO())

  err = cursor.All(context.TODO(), &resultsData)
  if err != nil {
    panic(err)
  }

  fmt.Println(resultsData)
}

func convertToStruct(csvData [][]string) []CsvData {
  data := []CsvData{}

  for indexRow, row := range csvData {
    if indexRow == 0 {
      continue
    }

    newCsvRowData := CsvData{}
    for indexCol, col := range row {
      switch indexCol {
      case 0:
        newCsvRowData.Id = col

      case 1:
        newCsvRowData.Name = col

      case 2:
        newCsvRowData.Email = col

      case 3:
        formatedAge, err := strconv.Atoi(col)
        if err != nil {
          log.Fatal(err)
        }
        newCsvRowData.Age = formatedAge

      case 4:
        formatedDateTime, err := convertStringToTime(col)
        if err != nil {
          log.Fatal(err)
        }
        newCsvRowData.Date = formatedDateTime
    }
    }

    data = append(data, newCsvRowData)
  }

  return data
}

func convertStringToTime(date string) (time.Time, error) {
  formatedDateTime, err := time.Parse("02/01/2006", date)
  if err != nil {
    return time.Time{}, err
  }

  return formatedDateTime, nil
}

func jsonEncode(data []CsvData) {
  json, err := json.MarshalIndent(data, "", "  ")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(json))
}

