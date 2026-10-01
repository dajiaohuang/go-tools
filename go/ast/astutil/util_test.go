package astutil

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestEqualCompositeLitElements(t *testing.T) {
	a := &ast.CompositeLit{Elts: []ast.Expr{
		&ast.BasicLit{Kind: token.INT, Value: "1"},
	}}
	b := &ast.CompositeLit{Elts: []ast.Expr{
		&ast.BasicLit{Kind: token.INT, Value: "2"},
	}}
	if Equal(a, b) {
		t.Fatal("Equal considered composite literals with different elements equal")
	}
}
