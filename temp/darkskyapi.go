package temp

import (
	"bytes"
	"fmt"
	"github.com/chr1sto14/nolatemp/net"
	"io/ioutil"
	"os/user"
)

const absoluteZeroF = 459.67

var nolaUrlRoot string = "https://api.darksky.net/forecast/%s/29.953,-90.071"
var ApiKey string

type DarkSkyApi struct {
	Currently WeatherData `json:"currently"`
}

type WeatherData struct {
	Temp float64 `json:"temperature"`
}

func getApiKey() (string, error) {
	// get user home dir
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	// grab api key
	datab, err := ioutil.ReadFile(usr.HomeDir + "/weatherapi.key")
	if err != nil {
		return "", err
	}

	return string(bytes.TrimSpace(datab)), nil
}

func getCurrentTemp(key string) (t float64, err error) {
	ds := new(DarkSkyApi)
	err = net.GetJson(fmt.Sprintf(nolaUrlRoot, key), ds)
	if err != nil {
		return
	}
	t = ds.Currently.Temp
	return
}

func GetNolaTemp() (temp float64, err error) {
	temp, err = getCurrentTemp(ApiKey)
	if err != nil {
		return
	}
	return
}
