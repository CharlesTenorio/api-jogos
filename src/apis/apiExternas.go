package apis

import (
	"api/src/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetTime(pagina int) {
	var ulrTime string = config.UrlApi["times"] + strconv.Itoa(pagina)

	response, err := http.Get(ulrTime)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(ulrTime))
	fmt.Println(string(responseData))

}
