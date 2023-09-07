package protocol

import (
	"fmt"
	"strings"
)

type msgVersionAck struct {
	header  msgHeader
	payload struct{}
}

func NewMsgVersionAck() *msgVersionAck {
	m := &msgVersionAck{
		header: msgHeader{
			magic:   MainNet,
			command: cmdVersionAck,
			length:  0,
		},
		payload: struct{}{},
	}
	m.header.checksum = msgChecksum(m)
	return m
}

func (m *msgVersionAck) String() string {
	var b strings.Builder
	fmt.Fprintln(&b, "Message Version Acknowledge")
	fmt.Fprintf(&b, "%s", m.header)
	fmt.Fprintf(&b, "  No Payload\n")
	return b.String()
}

func (m *msgVersionAck) pLoad() any {
	return m.payload
}
