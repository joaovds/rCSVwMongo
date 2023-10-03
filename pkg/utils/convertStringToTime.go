package utils

import "time"

func ConvertStringToTime(date string) (time.Time, error) {
  formatedDateTime, err := time.Parse("02/01/2006", date)
  if err != nil {
    return time.Time{}, err
  }

  return formatedDateTime, nil
}

