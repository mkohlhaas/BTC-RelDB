package protocol

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

type message interface {
	pLoad() any
}

func doubleSha256(data []byte) uint32 {
	var hash [32]byte
	hash = sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	return uint32(binary.BigEndian.Uint32(hash[:4]))
}

func msgChecksum(m message) checksum {
	buf := Encode(m.pLoad())
	return checksum(doubleSha256(buf))
}

func msgLength(m message) uint32 {
	buf := Encode(m.pLoad())
	return uint32(len(buf))
}

// Encode is the PUD
func Encode(m any) []byte {
	buf, err := Write(m)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Printf("%X\n", buf)
	return buf
}
