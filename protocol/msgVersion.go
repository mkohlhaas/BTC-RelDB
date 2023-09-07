package protocol

import (
	"fmt"
	"math/rand"
	"strings"
)

// TODO: separate files
type NetworkAddress struct{}
type VarString struct{}

type msgVersion struct {
	header  msgHeader
	payload msgVersionPayload
}

type msgVersionPayload struct {
	version     int32
	services    ServiceFlag
	timestamp   Timestamp
	addrTo      NetworkAddress
	addrFrom    NetworkAddress
	nonce       uint64
	userAgent   VarString
	startHeight int32
	relay       bool
}

func (m *msgVersion) pLoad() any {
	return m.payload
}

func (m *msgVersion) String() string {
	var b strings.Builder
	fmt.Fprintln(&b, "Message Version")
	fmt.Fprintf(&b, "%s", m.header)
	fmt.Fprintf(&b, "%s", m.payload)
	return b.String()
}

func (p msgVersionPayload) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "  Payload\n")
	fmt.Fprintf(&b, "    %-14s%d\n", "Version:", p.version)
	fmt.Fprintf(&b, "    %-14s%s\n", "Services:", p.services)
	fmt.Fprintf(&b, "    %-14s%s\n", "Timestamp:", p.timestamp)
	fmt.Fprintf(&b, "    %-14s%s\n", "Address To:", p.addrTo)
	fmt.Fprintf(&b, "    %-14s%s\n", "Address From:", p.addrFrom)
	fmt.Fprintf(&b, "    %-14s%d\n", "Nonce:", p.nonce)
	fmt.Fprintf(&b, "    %-14s%s\n", "User Agent:", p.userAgent)
	fmt.Fprintf(&b, "    %-14s%d\n", "Start Height:", p.startHeight)
	fmt.Fprintf(&b, "    %-14s%t\n", "Relay:", p.relay)
	return b.String()
}

func NewMsgVersion() *msgVersion {
	m := &msgVersion{
		header: msgHeader{
			magic:   MainNet,
			command: cmdVersion,
		},
		payload: msgVersionPayload{
			version:     ProtocolVersion,
			services:    SFNetwork,
			timestamp:   Now(),
			addrTo:      NetworkAddress{}, // TODO
			addrFrom:    NetworkAddress{}, // TODO
			nonce:       rand.Uint64(),    // TODO:save nonce
			userAgent:   VarString{},      // TODO
			startHeight: 0,
			relay:       false,
		},
	}
	m.header.length = msgLength(m)
	m.header.checksum = msgChecksum(m)
	return m
}
