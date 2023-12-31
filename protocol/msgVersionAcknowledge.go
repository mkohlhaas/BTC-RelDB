package protocol

import (
	"fmt"
	"strings"
)

type emtpyPayload struct{}

// MsgVersionAck = Message Version Acknowledge
type MsgVersionAck struct {
	Header  MsgHeader
	payload emtpyPayload
}

// NewMsgVersionAck creates a new version achnowledge message.
func NewMsgVersionAck() *MsgVersionAck {
	m := &MsgVersionAck{
		Header: MsgHeader{
			Magic:   mainNet.magic,
			Command: cmdVersionAck,
			Length:  0,
		},
		payload: emtpyPayload{},
	}
	m.Header.Length, m.Header.Checksum = msgLenAndChecksum(m)
	return m
}

func (m *MsgVersionAck) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Message Version Acknowledge\n")
	fmt.Fprintf(&b, "%s", m.Header)
	return b.String()
}

func (m *MsgVersionAck) pLoad() any {
	return m.payload
}
