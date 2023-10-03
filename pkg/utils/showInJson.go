package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/joaovds/rCSVwMongo/internal/entities"
)

func ShowInJson(data []entities.User) {
  json, err := json.MarshalIndent(data, "", "  ")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(json))
}
