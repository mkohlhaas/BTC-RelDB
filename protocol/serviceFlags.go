package protocol

import "strings"

type ServiceFlags uint64

const (
	SfNetwork ServiceFlags = 1 << iota
	SfGetutxo
	SfBloom
	SfWitness
	SfXthin
	SfInvalid5
	SfCompactFilters
	SfInvalid6
	SfInvalid7
	SfSegwit2X
	SfNetworkLimited
)

func (s ServiceFlags) String() string {
	res := make([]string, 0)
	if s&SfNetwork != 0 {
		res = append(res, "Network")
	}
	if s&SfGetutxo != 0 {
		res = append(res, "GetUTXO")
	}
	if s&SfBloom != 0 {
		res = append(res, "Bloom")
	}
	if s&SfWitness != 0 {
		res = append(res, "Witness")
	}
	if s&SfXthin != 0 {
		res = append(res, "Xthin")
	}
	if s&SfCompactFilters != 0 {
		res = append(res, "Compact Filters")
	}
	if s&SfSegwit2X != 0 {
		res = append(res, "Segwit2X")
	}
	if s&SfNetworkLimited != 0 {
		res = append(res, "Network Limited")
	}
	return strings.Join(res, " | ")
}
