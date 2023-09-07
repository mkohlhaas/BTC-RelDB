package main

import (
	"fmt"

	p "github.com/mkohlhaas/btcreldb/protocol"
)

func main() {
	mva := p.NewMsgVersionAck()
	fmt.Println(mva)
	fmt.Printf("On the wire: %X\n\n", p.OnTheWire(mva))
	mv := p.NewMsgVersion()
	fmt.Println(mv)
	fmt.Printf("On the wire: %X\n", p.OnTheWire(mv))
}
