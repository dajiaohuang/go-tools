package astutil

import (
	"go/ast"
	"go/parser"
	"testing"
)

func TestEqualCompositeLiterals(t *testing.T) {
	parse := func(src string) ast.Expr {
		expr, err := parser.ParseExpr(src)
		if err != nil {
			t.Fatal(err)
		}
		return expr
	}

	a := parse("[]int{1, 2}")
	if !Equal(a, parse("[]int{1, 2}")) {
		t.Fatal("Equal reports identical composite literals as different")
	}
	if Equal(a, parse("[]int{1, 3}")) {
		t.Fatal("Equal reports different composite literal elements as equal")
	}
}

func TestEqualChannelTypes(t *testing.T) {
	parse := func(src string) ast.Expr {
		expr, err := parser.ParseExpr(src)
		if err != nil {
			t.Fatal(err)
		}
		return expr
	}
	if Equal(parse("chan int"), parse("chan string")) {
		t.Fatal("Equal reports channel types with different element types as equal")
	}
}
