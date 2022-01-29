package ec2ifupifdown

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

var (
	macsURL           = "http://169.254.169.254/latest/meta-data/network/interfaces/macs"
	deviceNumberIdURL = "http://169.254.169.254/latest/meta-data/network/interfaces/macs/%s/device-number"
)

// GetInterfaces returns a list of al interfaces' mac addresses
func GetInterfaces() (interfaces []string, err error) {
	resp, err := http.Get(macsURL)
	if err != nil {
		return interfaces, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return interfaces, err
	}

	interfaces = strings.Split(string(body), "\n")

	if len(interfaces) == 0 {
		return interfaces, errors.New("no interfaces")
	}

	return interfaces, nil
}

// GetDeviceNumber returns the device number given the mac address
// ie: 0a:29:96:af:20:72 -> 1
func GetDeviceNumber(mac string) (deviceNumber int, err error) {
	url := fmt.Sprintf(deviceNumberIdURL, mac)

	resp, err := http.Get(url)
	if err != nil {
		return deviceNumber, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return deviceNumber, err
	}

	return strconv.Atoi(string(body))
}

// Ifdown disables interface
// devicenumber 0 -> eth0
// devicenumber 1 -> eth1
func Ifdown(deviceNumber int) (err error) {
	interf := fmt.Sprintf("eth%d", deviceNumber)

	cmd := exec.Command("ifdown", interf)
	return cmd.Run()
}

// Ifup enables interface
// devicenumber 0 -> eth0
// devicenumber 1 -> eth1
func Ifup(deviceNumber int) (err error) {
	interf := fmt.Sprintf("eth%d", deviceNumber)

	cmd := exec.Command("ifup", interf)
	return cmd.Run()
}
