package entities

import "time"

type User struct {
  CsvId    string    `json:"csvID" bson:"csvID"`
  Name     string    `json:"name"`
  Email    string    `json:"email"`
  Age      int       `json:"age"`
  Date     time.Time `json:"date"`
}

