package buildid

import (
	"bytes"
	"math"
	"testing"
)

func TestAligned4Size(t *testing.T) {
	for _, tc := range []struct {
		size int32
		want uint64
	}{
		{0, 0},
		{1, 4},
		{4, 4},
		{math.MaxInt32, 1 << 31},
	} {
		got, err := aligned4Size(tc.size)
		if err != nil || got != tc.want {
			t.Errorf("aligned4Size(%d) = %d, %v; want %d, nil", tc.size, got, err, tc.want)
		}
	}
	if _, err := aligned4Size(-1); err == nil {
		t.Error("aligned4Size(-1) succeeded")
	}
}

func TestReadAligned4(t *testing.T) {
	got, err := readAligned4(bytes.NewReader([]byte("abcx")), 3)
	if err != nil || string(got) != "abc" {
		t.Fatalf("readAligned4 = %q, %v; want abc, nil", got, err)
	}
	if _, err := readAligned4(bytes.NewReader([]byte("abc")), 3); err == nil {
		t.Fatal("readAligned4 accepted a truncated padded field")
	}
}
