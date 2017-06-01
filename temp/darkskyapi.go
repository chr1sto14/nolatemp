package temp

import (
	"bytes"
	"fmt"
	"github.com/chr1sto14/nolatemp/net"
	"io/ioutil"
	"log"
	"os/user"
	"strconv"
)

const absoluteZeroF = 459.67

var nolaUrlRoot string = "https://api.darksky.net/forecast/%s/29.953,-90.071"

type DarkSkyApi struct {
	Currently WeatherData `json:"currently"`
}

type WeatherData struct {
	Temp float64 `json:"temperature"`
}

func getCurrentWeather() (float64, error) {
	// TODO err check here
	user, _ := user.Current()

	// TODO err check here
	datab, _ := ioutil.ReadFile(user.HomeDir + "/weatherapi.key")
	data := bytes.TrimSpace(datab)
	// TODO fmt.Sprintf
	url := fmt.Sprintf(nolaUrlRoot, string(data))

	log.Printf("url %s", url)
	ds := new(DarkSkyApi)
	err := net.GetJson(url, ds)
	if err != nil {
		log.Printf("Error %s", err)
	}
	return ds.Currently.Temp, nil
}

func kStrToF(valStr string) (float64, error) {
	// TODO err check here
	val, _ := strconv.ParseFloat(valStr, 64)
	return float64((val * 9 / 5) - absoluteZeroF), nil
}

func GetNolaTemp() float64 {
	// TODO err check here
	temp, _ := getCurrentWeather()
	// TODO err check here
	// temp, _ := kStrToF(tempStr)
	return temp
}
