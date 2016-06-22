package main

import (
	"path/filepath"
	"testing"
)

func TestNumToFilename(t *testing.T) {
	testCases := []struct {
		val      int
		filename string
	}{
		{val: 0, filename: "0.b"},
		{val: 1, filename: "1.b"},
		{val: 9, filename: "9.b"},
		{val: 10, filename: "10.b"},
		{val: 99, filename: "99.b"},
		{val: 100, filename: "100.b"},
		{val: 999, filename: "999.b"},
		{val: 1000, filename: filepath.Join("1", "000.b")},
		{val: 1001, filename: filepath.Join("1", "001.b")},
		{val: 1999, filename: filepath.Join("1", "999.b")},
		{val: 2000, filename: filepath.Join("2", "000.b")},
		{val: 9999, filename: filepath.Join("9", "999.b")},
		{val: 10000, filename: filepath.Join("10", "000.b")},
		{val: 999999, filename: filepath.Join("999", "999.b")},
		{val: 1000000, filename: filepath.Join("1", "000", "000.b")},
	}
	for _, c := range testCases {
		got := numToFilename(c.val)
		if got != c.filename {
			t.Errorf("unexpected result for value=%d, got:%s; want:%s", c.val, got, c.filename)
		}
	}
}
