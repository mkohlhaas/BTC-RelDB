package protocol

import (
	"fmt"
	"strings"
)

type msgHeader struct {
	magic    btcNet
	command  [12]byte
	length   uint32
	checksum uint32
}

func (m msgHeader) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "  Message Header\n")
	fmt.Fprintf(&b, "    %-18s%s\n", "Bitcoin Network:", m.magic)
	fmt.Fprintf(&b, "    %-18s%s\n", "Command:", m.command)
	fmt.Fprintf(&b, "    %-18s%d\n", "Payload Length:", m.length)
	fmt.Fprintf(&b, "    %-18s%X\n", "Payload Checksum:", m.checksum)
	return b.String()
}
