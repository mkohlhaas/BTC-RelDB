package protocol

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
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

func msgLenAndChecksum(m message) (uint32, checksum) {
	buf := Encode(m.pLoad())
	return uint32(len(buf)), checksum(doubleSha256(buf))
}

// Encode is the PUD
func Encode(m any) []byte {
	buf, err := Write(m)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// fmt.Printf("%X\n", buf)
	return buf
}

// Encode is the PUD
func Decode(data []byte, m any) error {
	err := Read(data, m)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// fmt.Printf("%X\n", m)
	return nil
}

func Chk(err error) {
	if err != nil {
		switch err.Error() {
		case "EOF":
			fmt.Println("Client hung up.")
			os.Exit(0)
		case "ErrUnexpectedEOF":
			fmt.Println("Client hung up and we got too little data.")
			os.Exit(1)
		default:
			panic(err)
		}
	}
}

func TakeWhile[T any](s []T, predicate func(c T) bool) []T {
	result := make([]T, 0)
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		} else {
			break
		}
	}
	return result
}

func Filter[T any](s []T, predicate func(c T) bool) []T {
	result := make([]T, 0)
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func RemoveTrailingZeros(s []byte) []byte {
	return TakeWhile(s, func(c byte) bool {
		return c != 0x0
	})
}
