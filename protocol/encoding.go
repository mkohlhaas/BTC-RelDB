package protocol

import (
	"errors"
	"fmt"
	"reflect"
)

// Decode interface for all messages
// type Decode interface {
// 	decode(io.Reader)
// }

type coder struct {
	buf    []byte
	offset uint64
}

type decoder coder
type encoder coder

func sizeof(v reflect.Value) (sum int, err error) {
	t := v.Type()
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		sum = v.Len()
	case reflect.Struct:
		for i, n := 0, t.NumField(); i < n; i++ {
			s, err := sizeof(v.Field(i))
			if err != nil {
				return -1, fmt.Errorf("could not determine size of value ")
			}
			sum += s
		}
	case reflect.Uint64:
		switch t.Name() {
		case "VarInt":
			v1 := v.Uint()
			switch {
			case v1 <= 0xFC:
				sum++
			case v1 <= 0xFFFF:
				sum += 3
			case v1 <= 0xFFFFFFFF:
				sum += 5
			case v1 <= 0xFFFFFFFFFFFFFFFF:
				sum += 9
			}
		default:
			sum = int(t.Size())
		}
	default:
		sum = int(t.Size())
	}
	return
}

func (e *encoder) value(v reflect.Value) {
	t := v.Type()
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		l := v.Len()
		switch t.Name() {
		case "IP": // IP is big-endian (so-called network order)
			for i := l - 1; i >= 0; i-- {
				e.value(v.Index(i))
			}
		default: // is little-endian
			for i := 0; i < l; i++ {
				e.value(v.Index(i))
			}
		}
	case reflect.Struct:
		t := v.Type()
		l := v.NumField()
		for i := 0; i < l; i++ {
			if v := v.Field(i); v.CanSet() || t.Field(i).Name != "_" {
				e.value(v)
			}
		}
	case reflect.Bool:
		e.PutBool(v)
	case reflect.Uint8:
		e.PutUInt8(v)
	case reflect.Uint16:
		switch t.Name() {
		case "port": // port is big-endian (so-called network order)
			e.PutPort(v)
		default: // is little-endian
			e.PutUInt16(v)
		}
	case reflect.Int32:
		e.PutInt32(v)
	case reflect.Uint32:
		switch t.Name() {
		case "checksum": // checksum is big-endian (so-called internal byte order)
			e.PutChecksum(v)
		default:
			e.PutUInt32(v)
		}
	case reflect.Int64:
		e.PutInt64(v)
	case reflect.Uint64:
		switch t.Name() {
		case "VarInt":
			val := v.Uint()
			switch {
			case val <= 0xFC:
				e.PutUInt8(v)
			case val <= 0xFFFF:
				e.buf[e.offset] = byte(0xFD)
				e.offset++
				e.PutUInt16(v)
			case val <= 0xFFFFFFFF:
				e.buf[e.offset] = byte(0xFE)
				e.offset++
				e.PutUInt32(v)
			case val <= 0xFFFFFFFFFFFFFFFF:
				e.buf[e.offset] = byte(0xFF)
				e.offset++
				e.PutUInt64(v)
			}
		default:
			e.PutUInt64(v)
		}
	case reflect.String:
		e.PutString(v.String())
	default:
		fmt.Println("not implemented: ", t.Name())
	}
}

func (e *encoder) PutString(s string) {
	n := uint64(len(s))
	copy(e.buf[e.offset:e.offset+n], s)
	e.offset += n
}

func (e *encoder) PutBool(v reflect.Value) {
	val := v.Bool()
	if val {
		e.buf[e.offset] = 1
	} else {
		e.buf[e.offset] = 0
	}
	e.offset++
}

func (e *encoder) PutUInt8(v reflect.Value) {
	val := v.Uint()
	e.buf[e.offset] = byte(val)
	e.offset++
}

func (e *encoder) PutPort(v reflect.Value) {
	val := v.Uint()
	e.buf[e.offset+0] = byte(val >> 8)
	e.buf[e.offset+1] = byte(val)
	e.offset += 2
}

func (e *encoder) PutUInt16(v reflect.Value) {
	val := v.Uint()
	e.buf[e.offset+0] = byte(val)
	e.buf[e.offset+1] = byte(val >> 8)
	e.offset += 2
}

func (e *encoder) PutInt32(v reflect.Value) {
	val := v.Int()
	e.buf[e.offset+0] = byte(val)
	e.buf[e.offset+1] = byte(val >> 8)
	e.buf[e.offset+2] = byte(val >> 16)
	e.buf[e.offset+3] = byte(val >> 24)
	e.offset += 4
}

func (e *encoder) PutChecksum(v reflect.Value) {
	val := v.Uint()
	e.buf[e.offset+0] = byte(val >> 24)
	e.buf[e.offset+1] = byte(val >> 16)
	e.buf[e.offset+2] = byte(val >> 8)
	e.buf[e.offset+3] = byte(val)
	e.offset += 4
}

func (e *encoder) PutUInt32(v reflect.Value) {
	val := v.Uint()
	e.buf[e.offset+0] = byte(val)
	e.buf[e.offset+1] = byte(val >> 8)
	e.buf[e.offset+2] = byte(val >> 16)
	e.buf[e.offset+3] = byte(val >> 24)
	e.offset += 4
}

func (e *encoder) PutInt64(v reflect.Value) {
	val := v.Int()
	e.buf[e.offset+0] = byte(val)
	e.buf[e.offset+1] = byte(val >> 8)
	e.buf[e.offset+2] = byte(val >> 16)
	e.buf[e.offset+3] = byte(val >> 24)
	e.buf[e.offset+4] = byte(val >> 32)
	e.buf[e.offset+5] = byte(val >> 40)
	e.buf[e.offset+6] = byte(val >> 48)
	e.buf[e.offset+7] = byte(val >> 56)
	e.offset += 8
}

func (e *encoder) PutUInt64(v reflect.Value) {
	val := v.Uint()
	e.buf[e.offset+0] = byte(val)
	e.buf[e.offset+1] = byte(val >> 8)
	e.buf[e.offset+2] = byte(val >> 16)
	e.buf[e.offset+3] = byte(val >> 24)
	e.buf[e.offset+4] = byte(val >> 32)
	e.buf[e.offset+5] = byte(val >> 40)
	e.buf[e.offset+6] = byte(val >> 48)
	e.buf[e.offset+7] = byte(val >> 56)
	e.offset += 8
}

var varint uint64

func (d *decoder) value(v reflect.Value) {
	t := v.Type()
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		l := v.Len()
		switch t.Name() {
		case "IP": // IP is big-endian (so-called network order)
			for i := l - 1; i >= 0; i-- {
				d.value(v.Index(i))
			}
		default: // is little-endian
			for i := 0; i < l; i++ {
				d.value(v.Index(i))
			}
		}
	case reflect.Struct:
		t := v.Type()
		l := v.NumField()
		for i := 0; i < l; i++ {
			if v := v.Field(i); v.CanSet() || t.Field(i).Name != "_" {
				d.value(v)
			}
		}
	case reflect.Bool:
		setBool(v, d.bool())
	case reflect.Uint8:
		setUint8(v, d.uint8())
	case reflect.Uint16:
		switch t.Name() {
		case "port": // port is big-endian (so-called network order)
			setUint16(v, d.uint16Be())
		default: // is little-endian
			setUint16(v, d.uint16())
		}
	case reflect.Int32:
		setInt32(v, d.int32())
	case reflect.Uint32:
		switch t.Name() {
		case "checksum": // checksum is big-endian (so-called internal byte order)
			setUint32(v, d.uint32Be())
		default:
			setUint32(v, d.uint32())
		}
	case reflect.Int64:
		setInt64(v, d.int64())
	case reflect.Uint64:
		if t.Name() == "VarInt" {
			x := uint8(d.buf[d.offset])
			d.offset++
			switch x {
			case 0xFD:
				varint = uint64(d.uint16())
				v.SetUint(varint)
			case 0xFE:
				varint = uint64(d.uint32())
				v.SetUint(varint)
			case 0xFF:
				varint = d.uint64()
				v.SetUint(varint)
			default:
				varint = uint64(x)
				v.SetUint(varint)
			}
		} else {
			setUint64(v, d.uint64())
		}
	case reflect.String:
		v.SetString(string(d.buf[d.offset : d.offset+varint]))
		d.offset += varint
	default:
		fmt.Println("not implemented: ", t.Name())
	}
}

func setBool(v reflect.Value, x bool) {
	v.SetBool(x)
}

func setUint8(v reflect.Value, x uint8) {
	v.SetUint(uint64(x))
}

func setUint16(v reflect.Value, x uint16) {
	v.SetUint(uint64(x))
}

func setUint32(v reflect.Value, x uint32) {
	v.SetUint(uint64(x))
}

func setInt32(v reflect.Value, x int32) {
	v.SetInt(int64(x))
}

func setInt64(v reflect.Value, x int64) {
	v.SetInt(int64(x))
}

func setUint64(v reflect.Value, x uint64) {
	v.SetUint(uint64(x))
}

func (d *decoder) bool() bool {
	x := d.buf[d.offset]
	d.offset++
	return x != 0
}

func (d *decoder) uint8() uint8 {
	x := d.buf[d.offset]
	d.offset++
	return x
}

func (d *decoder) uint16() uint16 {
	b := d.buf[d.offset : d.offset+2]
	d.offset += 2
	return uint16(b[0]) | uint16(b[1])<<8
}

func (d *decoder) uint16Be() uint16 {
	b := d.buf[d.offset : d.offset+2]
	d.offset += 2
	return uint16(b[1]) | uint16(b[0])<<8
}

func (d *decoder) uint32Be() uint32 {
	b := d.buf[d.offset : d.offset+4]
	d.offset += 4
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func (d *decoder) uint32() uint32 {
	b := d.buf[d.offset : d.offset+4]
	d.offset += 4
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func (d *decoder) uint64() uint64 {
	b := d.buf[d.offset : d.offset+8]
	d.offset += 8
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func (d *decoder) int16() int16 {
	return int16(d.uint16())
}

func (d *decoder) int32() int32 {
	return int32(d.uint32())
}

func (d *decoder) int64() int64 {
	return int64(d.uint64())
}

// Write out message.
// func Write(w io.Writer, msg message) error {
func Write(msg any) ([]byte, error) {
	v := reflect.Indirect(reflect.ValueOf(msg))
	size, err := sizeof(v)
	if err != nil {
		return nil, errors.New("size of some values could not be determined " + reflect.TypeOf(msg).String())
	}
	// fmt.Println("/////////// Size: ", size, " ", v.Type().Name(), " //////////////////////////////////\n")
	buf := make([]byte, size)
	e := &encoder{buf: buf}
	e.value(v)
	// fmt.Printf("%X\n", e.buf)
	// _, err := w.Write(buf)
	// return err
	return buf, nil
}

// Read reads in a Bitcoin message.
func Read(buf []byte, msg any) error {
	v := reflect.Indirect(reflect.ValueOf(msg))
	// size, err := sizeof(v)
	// if err != nil {
	// 	return errors.New("size of some values could not be determined " + reflect.TypeOf(msg).String())
	// }
	// fmt.Println("/////////// Read Size: ", size, " ", v.Type().Name(), " ///////////////////////////////\n")
	d := &decoder{buf: buf}
	d.value(v)
	return nil
}
