package bytesconv

import (
	"bytes"
	"testing"
)

func TestStrToBytesAndBack(t *testing.T) {
	strings := []string{
		"",
		"hello",
		"世界",
		"1234567890",
		"with spaces and symbols!@#",
	}

	for _, s := range strings {
		b := StrToBytes(s)
		if string(b) != s {
			t.Errorf("StrToBytes(%q) = %v, want %v", s, b, []byte(s))
		}
		// Round-trip
		s2 := BytesToStr(b)
		if s2 != s {
			t.Errorf("BytesToStr(StrToBytes(%q)) = %q, want %q", s, s2, s)
		}
	}
}

func TestBytesToStrAndBack(t *testing.T) {
	byteSlices := [][]byte{
		[]byte(""),
		[]byte("hello"),
		[]byte("世界"),
		[]byte("1234567890"),
		[]byte("with spaces and symbols!@#"),
	}

	for _, b := range byteSlices {
		s := BytesToStr(b)
		if s != string(b) {
			t.Errorf("BytesToStr(%v) = %q, want %q", b, s, string(b))
		}
		// Round-trip
		b2 := StrToBytes(s)
		if !bytes.Equal(b2, b) {
			t.Errorf("StrToBytes(BytesToStr(%v)) = %v, want %v", b, b2, b)
		}
	}
}

func TestStrToBytes_Empty(t *testing.T) {
	b := StrToBytes("")
	if len(b) != 0 {
		t.Errorf("StrToBytes(\"\") = %v, want []", b)
	}
}

func TestBytesToStr_Empty(t *testing.T) {
	s := BytesToStr([]byte{})
	if s != "" {
		t.Errorf("BytesToStr([]) = %q, want \"\"", s)
	}
}
