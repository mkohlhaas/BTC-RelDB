package protocol

import (
	"strings"
)

type ServiceFlag uint64

const (
	SFNetwork ServiceFlag = 1 << iota
	SFGetutxo
	SFBloom
	SFWitness
	SFXthin
	SFInvalid_5
	SFCompact_Filters
	SFInvalid_6
	SFInvalid_7
	SFSegwit2X
	SFNetwork_Limited
)

func (s ServiceFlag) String() string {
	res := make([]string, 0)
	if s&SFNetwork != 0 {
		res = append(res, "Network")
	}
	if s&SFGetutxo != 0 {
		res = append(res, "GetUTXO")
	}
	if s&SFBloom != 0 {
		res = append(res, "Bloom")
	}
	if s&SFWitness != 0 {
		res = append(res, "Witness")
	}
	if s&SFXthin != 0 {
		res = append(res, "Xthin")
	}
	if s&SFCompact_Filters != 0 {
		res = append(res, "Compact Filters")
	}
	if s&SFSegwit2X != 0 {
		res = append(res, "Segwit2X")
	}
	if s&SFNetwork_Limited != 0 {
		res = append(res, "Network Limited")
	}
	return strings.Join(res, " | ")
}
