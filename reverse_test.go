package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{"Hello World", "dlroW olleH"},
		{"Pankaj", "jaknaP"},
	}
	for _, tc := range testCases {
		rev, err := Reverse(tc.in)
		if err != nil {
			return
		}
		if rev != tc.want {
			t.Errorf("Reverse %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testCases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testCases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := Reverse(orig)
		if err != nil {
			return
		}
		doubleRev, err := Reverse(rev)
		if err != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid utf8 string %q", rev)
		}
	})
}
