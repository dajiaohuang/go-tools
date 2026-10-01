package lint

import (
	"strings"
	"testing"
)

func TestDocumentationOptionsFormatAsSeparateLines(t *testing.T) {
	doc := &Documentation{Title: "Title", Options: []string{"first", "second"}}
	got := doc.Format(true)
	want := "Options\n    first\n    second\n"
	if !strings.Contains(got, want) {
		t.Fatalf("formatted options do not appear on separate lines:\n%s", got)
	}
}
