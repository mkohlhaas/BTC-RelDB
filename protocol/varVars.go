package protocol

import (
	"fmt"
)

type varString struct {
	Varint varInt
	Str    string
}

func (v varString) String() string {
	return string(v.Str)
}

func toVarString(s string) (res varString) {
	return varString{
		Varint: varInt(len(s)),
		Str:    s,
	}
}

type varInt uint64

func (v *varInt) String() string {
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
