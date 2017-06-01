package temp

import (
	"fmt"
	"github.com/chr1sto14/nolatemp/net"
	"io/ioutil"
	"log"
	"os/user"
	"strconv"
)

const absoluteZeroF = 459.67

var nolaUrlRoot string = "http://api.openweathermap.org/data/2.5/weather?id=4335045&APPID="

type OwmApi struct {
	Main WeatherData `json:"main"`
}

type WeatherData struct {
	Temp string `json:"temp"`
}

func getCurrentWeather() (string, error) {
	// TODO err check here
	user, _ := user.Current()

	// TODO err check here
	data, _ := ioutil.ReadFile(user.HomeDir + "/weatherapi.key")
	// TODO fmt.Sprintf
	url := nolaUrlRoot + string(data)

	owmapi := new(OwmApi)
	err := net.GetJson(url, owmapi)
	if err != nil {
		log.Printf("Error %s", err)
	}
	fmt.Println("here " + owmapi.Main.Temp)
	return owmapi.Main.Temp, nil
}

func kStrToF(valStr string) (float64, error) {
	// TODO err check here
	val, _ := strconv.ParseFloat(valStr, 64)
	return float64((val * 9 / 5) - absoluteZeroF), nil
}

func GetNolaTemp() float64 {
	// TODO err check here
	tempStr, _ := getCurrentWeather()
	// TODO err check here
	temp, _ := kStrToF(tempStr)
	return temp
}
