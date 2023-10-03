package flags

import (
  "flag"

  "github.com/joaovds/rCSVwMongo/internal/entities"
  "github.com/joaovds/rCSVwMongo/pkg/utils"
)

type Flags struct {
  showCsvDataInJson *bool
  help *bool
}

func SetupFlags(data []entities.User) {
  flags := Flags{
    showCsvDataInJson: flag.Bool("show-csv-data-in-json", false, "Mostra os dados do CSV no formato JSON"),
    help: flag.Bool("help", false, "Mostra os comandos disponíveis"),
  }

  flag.Parse()

  handleFlags(flags, data)
}

func handleFlags(flags Flags, data []entities.User) {
  switch {
  case *flags.showCsvDataInJson:
    utils.ShowInJson(data)

  case *flags.help:
    flag.PrintDefaults()

  default:
    flag.PrintDefaults()
}
}

