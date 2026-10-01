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

func TestCopyExprCopiesTypeFields(t *testing.T) {
	for _, src := range []string{"struct{ Field int }", "interface{ Method(int) }"} {
		expr, err := parser.ParseExpr(src)
		if err != nil {
			t.Fatal(err)
		}
		copy, ok := CopyExpr(expr)
		if !ok {
			t.Fatalf("CopyExpr(%s) failed", src)
		}
		var originalFields, copiedFields *ast.FieldList
		switch expr := expr.(type) {
		case *ast.StructType:
			originalFields = expr.Fields
			copiedFields = copy.(*ast.StructType).Fields
		case *ast.InterfaceType:
			originalFields = expr.Methods
			copiedFields = copy.(*ast.InterfaceType).Methods
		}
		if originalFields == copiedFields || originalFields.List[0] == copiedFields.List[0] || originalFields.List[0].Type == copiedFields.List[0].Type {
			t.Fatalf("CopyExpr(%s) shares type field nodes", src)
		}
	}
}

func TestCopyExprCopiesIndexListIndices(t *testing.T) {
	expr, err := parser.ParseExpr("f[int, string]")
	if err != nil {
		t.Fatal(err)
	}
	original := expr.(*ast.IndexListExpr)
	copy, ok := CopyExpr(original)
	if !ok {
		t.Fatal("CopyExpr failed")
	}
	copied := copy.(*ast.IndexListExpr)
	if &original.Indices[0] == &copied.Indices[0] {
		t.Fatal("CopyExpr shares the index slice")
	}
	copied.Indices[0] = ast.NewIdent("bool")
	if original.Indices[0].(*ast.Ident).Name != "int" {
		t.Fatal("modifying the copy changed the original")
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

func TestEqualInterfaceMethodTypes(t *testing.T) {
	parse := func(src string) ast.Expr {
		expr, err := parser.ParseExpr(src)
		if err != nil {
			t.Fatal(err)
		}
		return expr
	}
	if Equal(parse("interface{ M(int) }"), parse("interface{ M(string) }")) {
		t.Fatal("Equal reports interface methods with different parameter types as equal")
	}
}
