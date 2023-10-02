package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/joaovds/rCSVwMongo/internal/database"
	"github.com/joaovds/rCSVwMongo/pkg/handleCSV"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CsvData struct {
  Id    string    `json:"csvID" bson:"csvID"`
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

  csvData, err := handleCSV.ReadCSV("./assets/csvTest.csv")
  if err != nil {
    log.Fatalln("Error reading CSV file:", err)
  }

  data := convertToStruct(csvData)

  jsonEncode(data)

  resultsData, err := database.GetMany("tests_rCSVwMONGO", bson.D{})
  if err != nil {
    log.Panicln("Error getting data from database:", err)
  }

  resultData, err := database.GetOne("tests_rCSVwMONGO", bson.D{primitive.E{Key: "email", Value: "teste@teste.com"}})
  if err != nil {
    log.Panicln("Error getting data from database:", err)
  }

  resultCreate, err := database.InsertOne("tests_rCSVwMONGO", CsvData{
    Id: "1",
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

