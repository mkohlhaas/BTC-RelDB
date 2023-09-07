package protocol

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

type message interface {
	pLoad() any
}

func doubleSha256(data []byte) uint32 {
	var hash [32]byte
	hash = sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	return binary.BigEndian.Uint32(hash[:4])
}

func OnTheWire(m message) string {
	return bufferMessage(m).String()
}

func bufferMessage(m message) *bytes.Buffer {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, m)
	return buf
}

func bufferPayload(m message) *bytes.Buffer {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, m.pLoad())
	return buf
}

func msgChecksum(m message) uint32 {
	buf := bufferPayload(m)
	return doubleSha256([]byte(buf.String()))
}

func msgLength(m message) uint32 {
	buf := bufferPayload(m)
	return uint32(len(buf.String()))
}
