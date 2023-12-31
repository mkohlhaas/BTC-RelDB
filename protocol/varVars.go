package protocol

import (
	"fmt"
)

type varString struct {
	Varint VarInt
	Str    string
}

func (v varString) String() string {
	return string(v.Str)
}

// VarString is Bitcoin's variable string definition.
func VarString(s string) (res varString) {
	return varString{
		Varint: VarInt(len(s)),
		Str:    s,
	}
}

// VarInt is Bitcoin's variable integer definition.
type VarInt uint64

func (v *VarInt) String() string {
	return fmt.Sprintf("%d", v)
}

// func toVarInt(data uint64) (res varInt) {
// 	buf := new(bytes.Buffer)
// 	switch {
// 	case data <= 0xFC:
// 		res = []byte{uint8(data)}
// 	case data <= 0xFFFF:
// 		binary.Write(buf, binary.LittleEndian, uint16(data))
// 		res = append([]byte{0xFD}, []byte(buf.String())...)
// 	case data <= 0xFFFFFFFF:
// 		binary.Write(buf, binary.LittleEndian, uint32(data))
// 		res = append([]byte{0xFE}, []byte(buf.String())...)
// 	case data <= 0xFFFFFFFFFFFFFFFF:
// 		binary.Write(buf, binary.LittleEndian, data)
// 		res = append([]byte{0xFF}, []byte(buf.String())...)
// 	}
// 	return
// }
