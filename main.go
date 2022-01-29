package ec2ifupifdown

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type Interf struct {
	DeviceNumber        int
	InterfaceId         string
	LocalHostname       string
	LocalIpv4s          []string
	Mac                 string
	OwnerId             string
	SecurityGroupIds    []string
	SecurityGroups      []string
	SubnetId            string
	SubnetIpv4CidrBlock string
	VpcId               string
	VpcIpv4CidrBlock    string
	VpcIpv4CidrBlocks   []string
}

var (
	macsURL = "http://169.254.169.254/latest/meta-data/network/interfaces/macs"
	baseURL = "http://169.254.169.254/latest/meta-data/network/interfaces/macs/%s/%s"
)

func GetInterfaces() (interfaces []Interf, err error) {
	macs, err := GetMacs()
	if err != nil {
		return interfaces, err
	}

	log.Println(macs)

	for _, mac := range macs {
		var interf Interf

		interf.Mac = mac

		interf.DeviceNumber, err = GetDeviceNumber(mac)
		if err != nil {
			return interfaces, err
		}

		interf.InterfaceId, err = GetInterfaceId(mac)
		if err != nil {
			return interfaces, err
		}

		interf.LocalHostname, err = GetLocalHostname(mac)
		if err != nil {
			return interfaces, err
		}

		interf.LocalIpv4s, err = GetLocalIpv4s(mac)
		if err != nil {
			return interfaces, err
		}

		interf.Mac = mac

		interf.OwnerId, err = GetOwnerId(mac)
		if err != nil {
			return interfaces, err
		}

		interf.SecurityGroupIds, err = GetSecurityGroupIds(mac)
		if err != nil {
			return interfaces, err
		}

		interf.SecurityGroups, err = GetSecurityGroups(mac)
		if err != nil {
			return interfaces, err
		}

		interf.SubnetId, err = GetSubnetId(mac)
		if err != nil {
			return interfaces, err
		}

		interf.SubnetIpv4CidrBlock, err = GetSubnetIpv4CidrBlock(mac)
		if err != nil {
			return interfaces, err
		}

		interf.VpcId, err = GetVpcId(mac)
		if err != nil {
			return interfaces, err
		}

		interf.VpcIpv4CidrBlock, err = GetVpcIpv4CidrBlock(mac)
		if err != nil {
			return interfaces, err
		}

		interf.VpcIpv4CidrBlocks, err = GetVpcIpv4CidrBlocks(mac)
		if err != nil {
			return interfaces, err
		}

		interfaces = append(interfaces, interf)
	}

	return interfaces, err
}

// GetMacs returns a list of al interfaces' mac addresses
func GetMacs() (macs []string, err error) {
	resp, err := http.Get(macsURL)
	if err != nil {
		return macs, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return macs, err
	}

	macs = strings.Split(string(body), "\n")

	if len(macs) == 0 {
		return macs, errors.New("no interfaces")
	}

	return macs, nil
}

// Ifdown disables interface
// devicenumber 0 -> disables eth0
// devicenumber 1 -> disables eth1
func (i Interf) Ifdown() (err error) {
	interf := fmt.Sprintf("eth%d", i.DeviceNumber)

	cmd := exec.Command("ifdown", interf)
	return cmd.Run()
}

// Ifup enables interface
// devicenumber 0 -> enables eth0
// devicenumber 1 -> enables eth1
func (i Interf) Ifup() (err error) {
	interf := fmt.Sprintf("eth%d", i.DeviceNumber)

	cmd := exec.Command("ifup", interf)
	return cmd.Run()
}
