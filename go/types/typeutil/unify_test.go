package typeutil

import (
	"go/types"
	"testing"
)

func TestUnifyFailurePreservesExistingBindings(t *testing.T) {
	pkg := types.NewPackage("p", "p")
	constraint := types.NewInterfaceType(nil, nil).Complete()
	param := types.NewTypeParam(types.NewTypeName(0, pkg, "T", nil), constraint)
	bindings := map[*types.TypeParam]types.Type{param: types.Typ[types.Int]}

	if Unify(param, types.Typ[types.Bool], bindings) {
		t.Fatal("Unify succeeded with a conflicting existing binding")
	}
	if got := bindings[param]; got != types.Typ[types.Int] {
		t.Fatalf("binding after failed Unify = %v, want int", got)
	}
}
