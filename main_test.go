package main

import "testing"

func TestNumDigits(t *testing.T) {
	testCases := []struct {
		val int
		num int
	}{
		{val: 0, num: 1},
		{val: 1, num: 1},
		{val: 9, num: 1},
		{val: 10, num: 2},
		{val: 99, num: 2},
		{val: 100, num: 3},
		{val: 999, num: 3},
		{val: 1000, num: 4},
	}
	for _, c := range testCases {
		got := numDigits(c.val)
		if got != c.num {
			t.Errorf("unexpected result for value=%d, got:%d; want:%d", c.val, got, c.num)
		}
	}
}

func TestMaxPowerToTenBelow(t *testing.T) {
	testCases := []struct {
		val int
		num int
	}{
		{val: 0, num: 1},
		{val: 1, num: 1},
		{val: 9, num: 1},
		{val: 10, num: 10},
		{val: 99, num: 10},
		{val: 100, num: 100},
		{val: 999, num: 100},
		{val: 1000, num: 1000},
	}
	for _, c := range testCases {
		got := maxPowerToTenBelow(c.val)
		if got != c.num {
			t.Errorf("unexpected result for value=%d, got:%d; want:%d", c.val, got, c.num)
		}
	}
}

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
		{val: 1000, filename: "1/000.b"},
		{val: 1001, filename: "1/001.b"},
		{val: 1999, filename: "1/999.b"},
		{val: 2000, filename: "2/000.b"},
		{val: 9999, filename: "9/999.b"},
		{val: 10000, filename: "10/000.b"},
		{val: 999999, filename: "999/999.b"},
		{val: 1000000, filename: "1/000/000.b"},
	}
	for _, c := range testCases {
		got := numToFilename(c.val)
		if got != c.filename {
			t.Errorf("unexpected result for value=%d, got:%s; want:%s", c.val, got, c.filename)
		}
	}
}
