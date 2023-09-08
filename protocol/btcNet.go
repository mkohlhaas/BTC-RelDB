// Package protocol for writing packages on the network.
package protocol

type magic uint32

type btcNet struct {
	magic   magic
	comment string
	port    port
}

func (m magic) String() string {
	var bitcoinNets = map[magic]string{
		mainNet.magic:    "Main Network",
		regtestNet.magic: "Regression/Test Network",
		test3Net.magic:   "Test3 Network",
		sigNet.magic:     "Sig Network",
	}
	return bitcoinNets[m]
}

var (
	mainNet = btcNet{
		magic:   0xD9B4BEF9,
		comment: "Main Network",
		port:    8333,
	}
	test3Net = btcNet{
		magic:   0x0709110B,
		comment: "Test3 Network",
		port:    18333,
	}
	regtestNet = btcNet{
		magic:   0xDAB5BFFA,
		comment: "Regression/Test Network",
		port:    18444,
	}
	sigNet = btcNet{
		magic:   0x40CF030A,
		comment: "Sig Network",
		port:    38333,
	}
)
