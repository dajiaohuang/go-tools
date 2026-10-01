package typeutil

import (
	"go/types"
	"testing"
)

func TestMapNilValue(t *testing.T) {
	key := types.Typ[types.Int]
	m := new(Map[*int])

	m.Set(key, nil)
	if got, ok := m.At(key); !ok || got != nil {
		t.Fatalf("At(%v) = (%v, %t), want (nil, true)", key, got, ok)
	}

	called := false
	m.Iterate(func(gotKey types.Type, gotValue *int) {
		called = true
		if gotKey != key || gotValue != nil {
			t.Errorf("Iterate got (%v, %v), want (%v, nil)", gotKey, gotValue, key)
		}
	})
	if !called {
		t.Fatal("Iterate did not visit the nil-valued entry")
	}
}

func TestMapStringUnwrapsValues(t *testing.T) {
	key := types.Typ[types.Int]
	m := new(Map[string])
	m.Set(key, "value")

	if got, want := m.String(), `{int: "value"}`; got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}
