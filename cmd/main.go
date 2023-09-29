package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type CsvData struct {
  Id    string    `json:"id"`
  Name  string    `json:"name"`
  Email string    `json:"email"`
  Age   int       `json:"age"`
  Date  time.Time `json:"date"`
}

func main() {
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

  fmt.Println(data)
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

