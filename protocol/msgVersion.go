package protocol

import (
	"fmt"
	"math/rand"
	"strings"
)

// MsgVersion = Message Version
type MsgVersion struct {
	Header  msgHeader
	Payload msgVersionPayload
}

type msgVersionPayload struct {
	Version     int32
	Services    ServiceFlags
	Timestamp   timestamp
	AddrYou     NetworkAddress
	AddrMe      NetworkAddress
	Nonce       uint64
	UserAgent   varString
	StartHeight int32
	Relay       bool
}

func (m *MsgVersion) pLoad() any {
	return m.Payload
}

func (m *MsgVersion) String() string {
	var b strings.Builder
	fmt.Fprintln(&b, "Message Version")
	fmt.Fprintf(&b, "%s", m.Header)
	fmt.Fprintf(&b, "%s", m.Payload)
	return b.String()
}

func (p msgVersionPayload) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "  Payload\n")
	fmt.Fprintf(&b, "    %-14s%d\n", "Version:", p.Version)
	fmt.Fprintf(&b, "    %-14s%s\n", "Services:", p.Services)
	fmt.Fprintf(&b, "    %-14s%s\n", "Timestamp:", p.Timestamp)
	fmt.Fprintf(&b, "    %-14s%s\n", "Address To:", p.AddrYou)
	fmt.Fprintf(&b, "    %-14s%s\n", "Address From:", p.AddrMe)
	fmt.Fprintf(&b, "    %-14s%X\n", "Nonce:", p.Nonce)
	fmt.Fprintf(&b, "    %-14s%s\n", "User Agent:", p.UserAgent)
	fmt.Fprintf(&b, "    %-14s%d\n", "Start Height:", p.StartHeight)
	fmt.Fprintf(&b, "    %-14s%t\n", "Relay:", p.Relay)
	return b.String()
}

// NewMsgVersion creates a new version message.
func NewMsgVersion(to IP) *MsgVersion {
	m := &MsgVersion{
		Header: msgHeader{
			Magic:   mainNet.magic,
			Command: cmdVersion,
		},
		Payload: msgVersionPayload{
			Version:   protocolVersion,
			Services:  SfNetwork,
			Timestamp: now(),
			AddrYou: NetworkAddress{
				Service: SfNetwork,
				Ip:      to,
				Port:    mainNet.port,
			},
			AddrMe: NetworkAddress{
				Service: SfNetwork,
				Ip:      myPublicIPAddress(),
				Port:    mainNet.port,
			},
			Nonce:       rand.Uint64(),
			UserAgent:   toVarString(userAgent),
			StartHeight: getLatestHeight(), // should probably be zero
			Relay:       true},
	}
	// TODO: encode payload only once
	m.Header.Length = msgLength(m)
	m.Header.Checksum = msgChecksum(m)
	return m
}
