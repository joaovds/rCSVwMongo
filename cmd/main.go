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
  csvData, err := reader.ReadAll()
  if err != nil {
    log.Fatal(err)
  }

  // normalizar os dados agora
  data := convertToStruct(csvData)

  fmt.Println(data)
}

func convertToStruct(csvData [][]string) []CsvData {
  data := []CsvData{}

  return data
}

