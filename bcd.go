package bcd

import (
	"encoding/hex"
	"fmt"
)

//lbcd adds necessary "0" at the left side of byte array, if number of bytes is odd
func ASCII2Lbcd(data []byte) []byte {
	if len(data)%2 != 0 {
		return bcd(append(data, "0"...))
	}
	return bcd(data)
}

//rbcd adds necessary "0" at the left side of byte array, if number of bytes is odd
func ASCII2Rbcd(data []byte) []byte {
	if len(data)%2 != 0 {
		return bcd(append([]byte("0"), data...))
	}
	return bcd(data)
}

// Encode numeric in ascii into bsd (be sure len(data) % 2 == 0)
func bcd(data []byte) []byte {
	out := make([]byte, len(data)/2+1)
	n, err := hex.Decode(out, data)
	if err != nil {
		panic(err.Error())
	}
	return out[:n]
}

func Lbcd2ASCII(data []byte) []byte {
	r := bcd2Ascii(data)
	if r[len(r)-1] == 48 {
		r = r[:len(r)-1]
	}
	return r
}

func Rbcd2ASCII(data []byte) []byte {
	r := bcd2Ascii(data)
	if r[0] == 48 {
		return r[1:]
	}
	return r
}

func bcd2Ascii(data []byte) []byte {
	out := make([]byte, len(data)*2)
	n := hex.Encode(out, data)
	return out[:n]
}

func Dec2Rbcd(dec int) []byte {
	return ASCII2Rbcd([]byte(fmt.Sprint(dec)))
}

func Dec2Lbcd(dec int) []byte {
	return ASCII2Lbcd([]byte(fmt.Sprint(dec)))
}