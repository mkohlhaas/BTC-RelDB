// Package test provides test functions.
package test

import (
	"encoding/hex"
	"fmt"

	p "github.com/mkohlhaas/btcreldb/protocol"
)

func Test1() {
	// verack is always the same
	// Message Version Acknowledge
	//   Message Header
	//     Bitcoin Network:  Main Network
	//     Command:          verack
	//     Payload Length:   0
	//     Payload Checksum: 5DF6E0E2

	m1 := p.NewMsgVersionAck()
	fmt.Println(m1)
	p.Encode(m1)

	m2 := p.NewMsgVersion(p.IPAddr("34.45.56.231"))
	e := p.Encode(m2)
	var m4 p.MsgVersion
	p.Read(e, &m4)
	fmt.Println(m2)
	fmt.Println(&m4)
	fmt.Println(m4.Header.Length)
	fmt.Printf("m2 == m4 ? %t\n", *m2 == m4)

	var m6 p.MsgVersion
	p.Decode(e, &m6)
	fmt.Println("Using Decode:", &m6)
	fmt.Printf("m2 == m6 ? %t\n", *m2 == m6)

	var m3 p.MsgVersionAck
	x, _ := hex.DecodeString("F9BEB4D976657261636B000000000000000000005DF6E0E2")
	p.Read(x, &m3)
	fmt.Println(&m3)

	var m5 p.MsgVersion
	x, _ = hex.DecodeString("F9BEB4D976657273696F6E000000000066000000FB09A3B48" +
		"01101000100000000000000B9E60165000000000100000000000000E7382D22FFFF0000" +
		"0000000000000000208D01000000000000005B03DF2EFFFF00000000000000000000208" +
		"D9EE4828EF2877C9C102F62746372656C64623A302E302E312F42520C0001")
	p.Read(x, &m5)
	fmt.Println(&m5)
	fmt.Println(m5.Header.Magic)
	fmt.Println(m5.Payload.Timestamp)
}
