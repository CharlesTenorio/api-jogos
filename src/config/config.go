package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	Dev                = true
	Token              = ""
)
var UrlApi = make(map[string]string)

func CarregarConfig() {

	var error error
	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Porta, error = strconv.Atoi(os.Getenv("API_PORTA"))
	if error != nil {
		Porta = 9000
	}

	Dev, error = strconv.ParseBool(os.Getenv("DEV"))
	if error != nil {
		Dev = false
	}

	if Dev == true {
		StringConexaoBanco = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", "localhost", "5432",
			"postgres", "linux123", "apostadb")

	} else {
		StringConexaoBanco = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"),
			os.Getenv("PGUSER"), os.Getenv("PASSWORD_DB"), os.Getenv("PGDATABASE"))

	}

	//Token = os.Getenv("TOKEN")

	UrlApi["times"] = "https://api.b365api.com/v1/team?token=74409-ok98vsO3htGHPn&sport_id=1&page="
	UrlApi["ligas"] = "https://api.b365api.com/v1/league?token=74409-ok98vsO3htGHPn&sport_id=1&page="
	UrlApi["upcomingEvents"] = "https://api.b365api.com/v2/events/upcoming?sport_id=1&token=74409-ok98vsO3htGHPn"
	UrlApi["endedEvents"] = "https://api.b365api.com/v2/events/ended?sport_id=1&token=74409-ok98vsO3htGHPn"
	UrlApi["enventsOddsSummarys"] = "https://api.b365api.com/v2/event/odds/summary?token=74409-ok98vsO3htGHPn&event_id="
	UrlApi["even=tsSearch"] = "https://api.b365api.com/v1/events/search?token=74409-ok98vsO3htGHPn&sport_id=1&home="

}
