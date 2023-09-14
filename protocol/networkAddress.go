package protocol

import (
	"fmt"
	"strconv"
	"strings"
)

type port uint16 // port is big-endian

func (p port) String() string {
	return strconv.Itoa(int(p))
}

// NetworkAddress
type NetworkAddress struct {
	Services ServiceFlags
	IP       IP
	Port     port
}

func (n NetworkAddress) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Network Address: [ Services: %s, IP Address: %s, Port: %s ]", n.Services, n.IP, n.Port)
	return b.String()
}

type NetworkID uint8

// NetworkAddress2
type NetworkAddress2 struct {
	Time      timestamp
	Services  ServiceFlags2
	NetworkID NetworkID
	Addr      [32]byte
	Port      port
}

const (
	IpV4  NetworkID = 1 + iota // 4 	IPv4 address (globally routed internet)
	IpV6                       // 16 	IPv6 address (globally routed internet)
	TorV2                      // 10 	Tor v2 hidden service address
	TorV3                      // 32 	Tor v3 hidden service address
	I2P                        // 32 	I2P overlay network address
	Cjdns                      // 16 	Cjdns overlay network address
)
