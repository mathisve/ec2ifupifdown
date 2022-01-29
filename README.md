# ec2ifupifdown

go package to control network interfaces on ec2 instances, among other things.

## Example
```go
interfaces, err := ec2ifupifdown.GetInterfaces()
if err != nil {
    log.Panic(err)
}

for _, interf := range interfaces {
    log.Println(interf)

    interf.Ifdown()
    log.Printf("interface %s is down", interf.InterfaceId)

    interf.Ifup()
    log.Printf("interface %s is up again", interf.InterfaceId)
}
```

## Interface struct
```go
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
```