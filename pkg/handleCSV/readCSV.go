package handleCSV

import (
	"encoding/csv"
	"os"
)

func ReadCSV(filePath string) ([][]string, error) {
  file, err := os.Open(filePath)
  if err != nil {
    return nil, err
  }

  defer file.Close()

  reader := csv.NewReader(file)
  reader.Comma = 59 // == ';'
  csvData, err := reader.ReadAll()
  if err != nil {
    return nil, err
  }

  return csvData, nil;
}

