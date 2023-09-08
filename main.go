// Package main is the main entry point for the application.
package main

import (
	p "github.com/mkohlhaas/btcreldb/protocol"
)

// TODO: ENCODING varInt and varString!!!

func main() {
	// m1 := p.NewMsgVersionAck()
	// fmt.Println(m1)
	// p.Encode(m1)

	// Message Version Acknowledge
	//   Message Header
	//     Bitcoin Network:  Main Network
	//     Command:          verack
	//     Payload Length:   0
	//     Payload Checksum: 5DF6E0E2

	m2 := p.NewMsgVersion(p.IPAddr("34.45.56.231"))
	// fmt.Println(m2)
	p.Encode(m2)

	// var m3 p.MsgVersionAck
	// x, _ := hex.DecodeString("F9BEB4D976657261636B000000000000000000005DF6E0E2")
	// p.Read(&m3, x)
	// fmt.Println(m3)

	// var m4 p.MsgVersion
	// x, _ := hex.DecodeString("F9BEB4D976657273696F6E000000000066000000FB09A3B4801101000100000000000000B9E60165000000000100000000000000E7382D22FFFF00000000000000000000208D01000000000000005B03DF2EFFFF00000000000000000000208D9EE4828EF2877C9C102F62746372656C64623A302E302E312F42520C0001")
	// p.Read(x, &m4)
	// fmt.Println(m4)
	// fmt.Println(m4.Header.Magic)
	// fmt.Println(m4.Payload.Timestamp)
}
