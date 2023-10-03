package handleCSV

import (
	"log"
	"strconv"

	"github.com/joaovds/rCSVwMongo/internal/entities"
	"github.com/joaovds/rCSVwMongo/pkg/utils"
)

func ConvertToStruct(csvData [][]string) []entities.User {
  data := []entities.User{}

  for indexRow, row := range csvData {
    if indexRow == 0 {
      continue
    }

    newCsvRowData := entities.User{}
    for indexCol, col := range row {
      switch indexCol {
      case 0:
        newCsvRowData.CsvId = col

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
        formatedDateTime, err := utils.ConvertStringToTime(col)
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

