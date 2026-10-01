package nilness

import (
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNilness(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analysis, "example.com/...")
}

func TestNilnessMismatchedFactArity(t *testing.T) {
	results := types.NewTuple(
		types.NewVar(token.NoPos, nil, "", types.Typ[types.Int]),
		types.NewVar(token.NoPos, nil, "", types.Typ[types.Int]),
		types.NewVar(token.NoPos, nil, "", types.NewPointer(types.Typ[types.Int])),
	)
	sig := types.NewSignatureType(nil, nil, nil, types.NewTuple(), results, false)
	fn := types.NewFunc(token.NoPos, nil, "f", sig)
	r := &Result{m: map[*types.Func][]ValueNilness{
		fn: {{Outer: NeverNil}},
	}}

	got := r.Nilness(fn, 2)
	want := ValueNilness{Inner: MaybeNil, Outer: MaybeNil}
	if got != want {
		t.Fatalf("Nilness(fn, 2) = %v, want %v", got, want)
	}
}
