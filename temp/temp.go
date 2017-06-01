package temp

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var deviceFile string = "1000"

// getFile()

func getFile() string {
	baseDir := "/sys/bus/w1/devices/"
	devicesFolder, err := filepath.Glob(baseDir + "28*")
	if err != nil {
		return ""
	}
	return devicesFolder[0] + "/w1_slave"
}

func InitTemp() error {
	_, err := exec.Command("modprobe w1-gpio").Output()
	if err != nil {
		return err
	}

	_, err = exec.Command("modprobe w1-therm").Output()
	if err != nil {
		return err
	}
	return nil
}

func GetTemp() float64 {
	// TODO err check here
	data, _ := ioutil.ReadFile(deviceFile)

	datastr := string(data)

	bottomLine := strings.Split(datastr, "\n")[1]
	tempStr := strings.Split(bottomLine, " ")[9]
	// TODO err check here
	tempFloat, _ := strconv.ParseFloat(tempStr, 64)
	return tempFloat / 1000
}
