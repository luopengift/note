package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func HexToIPv4(s string) net.IP {
	ipList := make([]byte, 4)
	for i := range ipList {
		part, _ := strconv.ParseInt(s[i*2:i*2+2], 16, 0)
		ipList[4-1-i] = byte(part)
	}
	return net.IPv4(ipList[0], ipList[1], ipList[2], ipList[3])
}

func HexToIPv6(s string) net.IP {
	ipList := make([]string, 8)
	for i := range ipList {
		ipList[i] = s[i*4 : i*4+4]
	}
	return net.ParseIP(strings.Join(ipList, ":"))
}

func HexToIP(s string) net.IP {
	switch len(s) {
	case 8:
		return HexToIPv4(s)
	case 32:
		return HexToIPv6(s)
	default:
		return nil
	}
}

func main() {

	fmt.Println(HexToIP("70EB800A"))
	fmt.Println(HexToIP("0000000000000000FFFF00000100007F"))
	fmt.Println(HexToIP("1234"))
}
