package temp

import (
	"bytes"
	"fmt"
	"github.com/chr1sto14/nolatemp/net"
	"io/ioutil"
	"log"
	"os/user"
)

const absoluteZeroF = 459.67

var nolaUrlRoot string = "https://api.darksky.net/forecast/%s/29.953,-90.071"

type DarkSkyApi struct {
	Currently WeatherData `json:"currently"`
}

type WeatherData struct {
	Temp float64 `json:"temperature"`
}

func getApiKey() (string, error) {
	// TODO err check here
	usr, _ := user.Current()

	// TODO err check here
	datab, _ := ioutil.ReadFile(usr.HomeDir + "/weatherapi.key")

	return string(bytes.TrimSpace(datab)), nil
}

func getCurrentTemp(key string) (float64, error) {
	ds := new(DarkSkyApi)
	err := net.GetJson(fmt.Sprintf(nolaUrlRoot, key), ds)
	if err != nil {
		log.Printf("Error %s", err)
	}
	return ds.Currently.Temp, nil
}

// TODO err check here
func GetNolaTemp() float64 {
	key, _ := getApiKey()
	// TODO err check here
	temp, _ := getCurrentTemp(key)
	return temp
}
