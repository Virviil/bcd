package bcd

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBcd(t *testing.T) {
	Convey("Describe ASCII->BCD converter", t, func() {
		Convey("If BCD encoded value is", func() {
			Convey("valid, it should be converted right", func() {
				var x = []struct {
					value    string
					expected string
				}{
					{"00", "\x00"},
					{"90", "\x90"},
					{"1234", "\x12\x34"},
					{"AC", "\xAC"},
					{"ACD0", "\xAC\xD0"},
					{"acd0", "\xAC\xD0"},
					{"abcdef", "\xAB\xCD\xEF"},
				}
				for _, tc := range x {
					So(bcd([]byte(tc.value)), ShouldResemble, []byte(tc.expected))
				}
			})

			Convey("not valid, it should panic", func() {
				var x = []struct {
					value    string
					expected string
				}{
					{"qw", "encoding/hex: invalid byte: U+0071 'q'"},
					{"12345", "encoding/hex: odd length hex string"},
				}
				for _, tc := range x {
					So(func() { bcd([]byte(tc.value)) }, ShouldPanicWith, tc.expected)
				}
			})
		})
		Convey("If the BCD byte array is not valid - Rbcd should append 0 at the beginning of array", func() {
			var x = []struct {
				value    string
				expected string
			}{
				{"0", "00"},
				{"90", "90"},
				{"123", "0123"},
				{"AC", "AC"},
				{"ACD", "0ACD"},
			}
			for _, tc := range x {
				So(fmt.Sprintf("%X", ASCII2Rbcd([]byte(tc.value))), ShouldEqual, tc.expected)
			}
		})
		Convey("If the BCD byte array is not valid - Lbcd should append 0 at the end of array", func() {
			var x = []struct {
				value    string
				expected string
			}{
				{"0", "00"},
				{"90", "90"},
				{"123", "1230"},
				{"AC", "AC"},
				{"ACD", "ACD0"},
			}
			for _, tc := range x {
				So(fmt.Sprintf("%X", ASCII2Lbcd([]byte(tc.value))), ShouldEqual, tc.expected)
			}
		})
	})
	Convey("Describe BCD->ASCII converter", t, func() {
		Convey("It should convert byte data right with Rbcd", func() {
			var x = []struct {
				value    string
				expected string
			}{
				{"\x00", "0"},
				{"\x90\x40", "9040"},
				{"\x09\x04", "904"},
				{"\x12\x34", "1234"},
				{"\xAC", "ac"},
                {"\xac", "ac"},
				{"\xAC\xD0", "acd0"},
				{"\xac\xd0", "acd0"},
                {"\x0c\xd0", "cd0"},
				{"\xab\xcd\xef", "abcdef"},
			}
			for _, tc := range x {
				So(Rbcd2ASCII([]byte(tc.value)), ShouldResemble, []byte(tc.expected))
			}
		})
		Convey("It should convert byte data right with Lbcd", func() {
			var x = []struct {
				value    string
				expected string
			}{
				{"\x00", "0"},
				{"\x90\x40", "904"},
				{"\x09\x04", "0904"},
				{"\x12\x34", "1234"},
				{"\xAC", "ac"},
                {"\xac", "ac"},
				{"\xAC\xD0", "acd"},
				{"\xac\xd0", "acd"},
                {"\x0c\xd0", "0cd"},
				{"\xab\xcd\xef", "abcdef"},
			}
			for _, tc := range x {
				So(Lbcd2ASCII([]byte(tc.value)), ShouldResemble, []byte(tc.expected))
			}
		})
	})
}
