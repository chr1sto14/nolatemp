package temp

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var deviceFile string = getFile()

func getFile() string {
	baseDir := "/sys/bus/w1/devices/"
	devicesFolder, err := filepath.Glob(baseDir + "28*")
	if err != nil {
		return ""
	}
	return devicesFolder[0] + "/w1_slave"
}

func InitTemp() (err error) {
	if _, ok := os.LookupEnv("NOLATEMP"); ok {
		return
	} else {
		_, err = exec.Command("/sbin/modprobe w1-gpio").Output()
		if err != nil {
			return
		}

		_, err = exec.Command("/sbin/modprobe w1-therm").Output()
		if err != nil {
			return
		}
		os.Setenv("NOLATEMP", "true")
	}
	return
}

func GetTemp() (tempFloat float64, err error) {
	data, err := ioutil.ReadFile(deviceFile)
	if err != nil {
		return
	}

	bottomLine := strings.Split(string(data), "\n")[1]
	tempStr := strings.Split(bottomLine, " ")[9]
	tempFloat, err = strconv.ParseFloat(tempStr[2:], 64)
	if err != nil {
		return
	}
	tempFloat /= 1000
	return
}
