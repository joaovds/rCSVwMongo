package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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
    fmt.Printf("%T", row)

    // newCsvRowData := CsvData{}
    // for indexCol, col := range row {

    // }
  }

  return data
}

