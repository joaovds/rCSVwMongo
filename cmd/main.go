package main

import (
  "log"

  "github.com/joaovds/rCSVwMongo/internal/flags"
  "github.com/joaovds/rCSVwMongo/pkg/handleCSV"
  "github.com/joho/godotenv"
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

  dataFormated := handleCSV.ConvertToStruct(csvData)

  flags.SetupFlags(dataFormated)
}

