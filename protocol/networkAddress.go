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

type NetworkAddress struct {
	Service ServiceFlags
	Ip      IP
	Port    port
}

func (n NetworkAddress) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Network Address: [ Services: %s, IP Address: %s, Port: %s ]", n.Service, n.Ip, n.Port)
	return b.String()
}
