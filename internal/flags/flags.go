package flags

import (
	"flag"
	"fmt"
	"log"

	"github.com/joaovds/rCSVwMongo/internal/database"
	"github.com/joaovds/rCSVwMongo/internal/entities"
	"github.com/joaovds/rCSVwMongo/pkg/utils"
)

type Flags struct {
  showCsvDataInJson *bool
  saveDataInMongo *bool
  help *bool
}

func SetupFlags(data []entities.User) {
  flags := Flags{
    showCsvDataInJson: flag.Bool("show-csv-data-in-json", false, "Mostra os dados do CSV no formato JSON"),
    saveDataInMongo: flag.Bool("save-data-in-mongo", false, "Salva os dados do CSV no MongoDB"),
    help: flag.Bool("help", false, "Mostra os comandos disponíveis"),
  }

  flag.Parse()

  handleFlags(flags, data)
}

func handleFlags(flags Flags, data []entities.User) {
  switch {
  case *flags.showCsvDataInJson:
    utils.ShowInJson(data)

  case *flags.saveDataInMongo:
    for _, user := range data {
      success, err := database.InsertOne("tests_rCSVwMONGO", user)
      if err != nil {
        log.Panicln("Error inserting data in MongoDB:", err)
      }
      fmt.Println(success)
    }

  case *flags.help:
    flag.PrintDefaults()

  default:
    flag.PrintDefaults()
}
}

