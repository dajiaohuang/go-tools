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

func TestNilMapReadOperations(t *testing.T) {
	var m *Map[int]
	key := types.Typ[types.Int]

	if got, ok := m.At(key); ok || got != 0 {
		t.Fatalf("At(%v) = (%d, %t), want (0, false)", key, got, ok)
	}
	if m.Delete(key) {
		t.Fatal("Delete on nil map returned true")
	}
	if m.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", m.Len())
	}
	called := false
	m.Iterate(func(types.Type, int) { called = true })
	if called {
		t.Fatal("Iterate called its callback on a nil map")
	}
	if keys := m.Keys(); keys == nil || len(keys) != 0 {
		t.Fatalf("Keys() = %v, want a non-nil empty slice", keys)
	}
	if got := m.String(); got != "{}" {
		t.Fatalf("String() = %q, want {}", got)
	}
	if got := m.KeysString(); got != "{}" {
		t.Fatalf("KeysString() = %q, want {}", got)
	}
	m.SetHasher(MakeHasher())
}
