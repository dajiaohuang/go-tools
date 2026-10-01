package astutil

import (
	"go/ast"
	"testing"
)

func TestCopyExprRejectsShallowTypeCopies(t *testing.T) {
	for _, expr := range []ast.Expr{
		&ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "int"}}}}},
		&ast.InterfaceType{Methods: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "M"}}}}},
	} {
		if got, ok := CopyExpr(expr); ok || got != nil {
			t.Errorf("CopyExpr(%T) = (%T, %v), want (nil, false)", expr, got, ok)
		}
	}
}
