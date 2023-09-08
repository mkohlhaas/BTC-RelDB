package protocol

import (
	"fmt"
	"strings"
)

type checksum uint32 // checksum is big-endian

type msgHeader struct {
	Magic    magic
	Command  [12]byte
	Length   uint32
	Checksum checksum
}

func (m msgHeader) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "  Message Header\n")
	fmt.Fprintf(&b, "    %-18s%s\n", "Bitcoin Network:", m.Magic)
	fmt.Fprintf(&b, "    %-18s%s\n", "Command:", m.Command)
	fmt.Fprintf(&b, "    %-18s%d\n", "Payload Length:", m.Length)
	fmt.Fprintf(&b, "    %-18s%X\n", "Payload Checksum:", m.Checksum)
	return b.String()
}
