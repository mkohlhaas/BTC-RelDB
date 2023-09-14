package protocol

import (
	"fmt"
	"math/rand"
	"strings"
)

// MsgVersion = Message Version
type MsgVersion struct {
	Header  MsgHeader
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
		Header: MsgHeader{
			Magic:   mainNet.magic,
			Command: cmdVersion,
		},
		Payload: msgVersionPayload{
			Version:   protocolVersion,
			Services:  SfNetwork,
			Timestamp: now(),
			AddrYou: NetworkAddress{
				Services: SfNetwork,
				IP:       to,
				Port:     mainNet.port,
			},
			AddrMe: NetworkAddress{
				Services: SfNetwork,
				IP:       myPublicIPAddress(),
				Port:     mainNet.port,
			},
			Nonce:       rand.Uint64(),
			UserAgent:   VarString(userAgent),
			StartHeight: getLatestHeight(), // TODO: should probably be zero
			Relay:       false},
	}
	m.Header.Length, m.Header.Checksum = msgLenAndChecksum(m)
	return m
}
