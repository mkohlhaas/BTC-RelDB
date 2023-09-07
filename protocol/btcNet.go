package protocol

type btcNet uint32

const (
	MainNet     btcNet = 0xF9BEB4D9
	RegtestNet  btcNet = 0xFABFB5DA
	Test3Net    btcNet = 0x0B110907
	SigNet      btcNet = 0x0A03CF40
	NamecoinNet btcNet = 0xF9BEB4FE
)

func (b btcNet) String() string {
	var bitcoinNets = map[btcNet]string{
		MainNet:     "Main Network",
		RegtestNet:  "Regression/Test Network",
		Test3Net:    "Test3 Network",
		SigNet:      "Sig Network",
		NamecoinNet: "Namecoin Network",
	}
	return bitcoinNets[b]
}
