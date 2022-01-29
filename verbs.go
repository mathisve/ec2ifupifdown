package ec2ifupifdown

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	deviceNumberVerb        = "device-number"
	interfaceIdVerb         = "interface-id"
	localHostnameVerb       = "local-hostname"
	localIpv4sVerb          = "local-ipv4s"
	ownerIdVerb             = "owner-id"
	securityGroupIdsVerb    = "security-group-ids"
	securityGroupsVerb      = "security-groups"
	subnetIdVerb            = "subnet-id"
	subnetIpv4CidrBlockVerb = "subnet-ipv4-cidr-block"
	vpcIdVerb               = "vpc-id"
	vpcIpv4CidrBlockVerb    = "vpc-ipv4-cidr-block"
	vpcIpv4CidrBlocksVerb   = "vpc-ipv4-cidr-blocks"
)

func GetDeviceNumber(mac string) (deviceNumber int, err error) {
	url := fmt.Sprintf(baseURL, mac, deviceNumberVerb)

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

func GetInterfaceId(mac string) (interfaceId string, err error) {
	url := fmt.Sprintf(baseURL, mac, interfaceIdVerb)

	resp, err := http.Get(url)
	if err != nil {
		return interfaceId, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return interfaceId, err
	}

	return string(body), err
}

func GetLocalHostname(mac string) (localHostname string, err error) {
	url := fmt.Sprintf(baseURL, mac, localHostnameVerb)

	resp, err := http.Get(url)
	if err != nil {
		return localHostname, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return localHostname, err
	}

	return string(body), err
}

func GetLocalIpv4s(mac string) (localIpv4s []string, err error) {
	url := fmt.Sprintf(baseURL, mac, localIpv4sVerb)

	resp, err := http.Get(url)
	if err != nil {
		return localIpv4s, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return localIpv4s, err
	}

	return strings.Split(string(body), "\n"), err
}

func GetMac(mac string) (string, error) {
	return mac, nil
}

func GetOwnerId(mac string) (ownerId string, err error) {
	url := fmt.Sprintf(baseURL, mac, ownerIdVerb)

	resp, err := http.Get(url)
	if err != nil {
		return ownerId, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ownerId, err
	}

	return string(body), err
}

func GetSecurityGroupIds(mac string) (securityGroupIds []string, err error) {
	url := fmt.Sprintf(baseURL, mac, securityGroupIdsVerb)

	resp, err := http.Get(url)
	if err != nil {
		return securityGroupIds, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return securityGroupIds, err
	}

	return strings.Split(string(body), "\n"), err
}

func GetSecurityGroups(mac string) (securityGroups []string, err error) {
	url := fmt.Sprintf(baseURL, mac, securityGroupsVerb)

	resp, err := http.Get(url)
	if err != nil {
		return securityGroups, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return securityGroups, err
	}

	return strings.Split(string(body), "\n"), err
}

func GetSubnetId(mac string) (subnetId string, err error) {
	url := fmt.Sprintf(baseURL, mac, subnetIdVerb)

	resp, err := http.Get(url)
	if err != nil {
		return subnetId, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return subnetId, err
	}

	return string(body), err
}

func GetSubnetIpv4CidrBlock(mac string) (subnetIpv4CidrBlock string, err error) {
	url := fmt.Sprintf(baseURL, mac, subnetIpv4CidrBlockVerb)

	resp, err := http.Get(url)
	if err != nil {
		return subnetIpv4CidrBlock, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return subnetIpv4CidrBlock, err
	}

	return string(body), err
}

func GetVpcId(mac string) (vpcId string, err error) {
	url := fmt.Sprintf(baseURL, mac, vpcIdVerb)

	resp, err := http.Get(url)
	if err != nil {
		return vpcId, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return vpcId, err
	}

	return string(body), err
}

func GetVpcIpv4CidrBlock(mac string) (vpcIpv4CidrBlock string, err error) {
	url := fmt.Sprintf(baseURL, mac, vpcIpv4CidrBlockVerb)

	resp, err := http.Get(url)
	if err != nil {
		return vpcIpv4CidrBlock, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return vpcIpv4CidrBlock, err
	}

	return string(body), err
}

func GetVpcIpv4CidrBlocks(mac string) (vpcIpv4CidrBlocks []string, err error) {
	url := fmt.Sprintf(baseURL, mac, vpcIpv4CidrBlocksVerb)

	resp, err := http.Get(url)
	if err != nil {
		return vpcIpv4CidrBlocks, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return vpcIpv4CidrBlocks, err
	}

	return strings.Split(string(body), "\n"), err
}
