package temp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const absoluteZeroF = 459.67

var nolaUrlRoot string = "https://api.openweathermap.org/data/2.5/weather?zip=70112,us&APPID="

func parseWeather(wInfo []byte) (tempStr string, err error) {
	// TODO err check here
	var blob map[string]interface{}
	err = json.Unmarshal(wInfo, &blob)

	// TODO ok check here
	mainBlob, _ := blob["main"].(map[string]interface{})

	// TODO ok check here
	tempStr, _ = mainBlob["temp"].(string)
	return
}

func getCurrentWeather() (string, error) {
	// TODO err check here
	data, _ := ioutil.ReadFile("~/weatherapi.key")
	url := nolaUrlRoot + string(data)

	// TODO err check here
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	// TODO err check here
	wInfo, _ := ioutil.ReadAll(resp.Body)

	// TODO err check here
	tempStr, _ := parseWeather(wInfo)

	return tempStr, nil
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
