package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

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

  fmt.Println(csvData)
}

