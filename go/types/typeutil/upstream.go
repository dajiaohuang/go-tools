package typeutil

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"
	_ "unsafe"

	"golang.org/x/tools/go/types/typeutil"
)

type MethodSetCache = typeutil.MethodSetCache
type Hasher = typeutil.Hasher

func Callee(info *types.Info, call *ast.CallExpr) types.Object {
	return typeutil.Callee(info, call)
}

func IntuitiveMethodSet(T types.Type, msets *MethodSetCache) []*types.Selection {
	return typeutil.IntuitiveMethodSet(T, msets)
}

func MakeHasher() Hasher {
	return typeutil.MakeHasher()
}

type Map[V any] struct {
	m typeutil.Map
}

type mapValue[V any] struct{ value V }

func (m *Map[V]) Delete(key types.Type) bool { return m.m.Delete(key) }
func (m *Map[V]) At(key types.Type) (V, bool) {
	v := m.m.At(key)
	if v == nil {
		var zero V
		return zero, false
	}
	return v.(mapValue[V]).value, true
}
func (m *Map[V]) Set(key types.Type, value V) { m.m.Set(key, mapValue[V]{value}) }
func (m *Map[V]) Len() int                    { return m.m.Len() }
func (m *Map[V]) Iterate(f func(key types.Type, value V)) {
	ff := func(key types.Type, value any) {
		f(key, value.(mapValue[V]).value)
	}
	m.m.Iterate(ff)

}
func (m *Map[V]) Keys() []types.Type          { return m.m.Keys() }
func (m *Map[V]) String() string {
	var b strings.Builder
	b.WriteByte('{')
	sep := ""
	m.Iterate(func(key types.Type, value V) {
		fmt.Fprintf(&b, "%s%v: %q", sep, key, any(value))
		sep = ", "
	})
	b.WriteByte('}')
	return b.String()
}
func (m *Map[V]) KeysString() string          { return m.m.KeysString() }
func (m *Map[V]) SetHasher(h typeutil.Hasher) { m.m.SetHasher(h) }
