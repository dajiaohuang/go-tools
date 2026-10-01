package lint

import (
	"reflect"
	"testing"
)

func TestParseDirectiveWhitespace(t *testing.T) {
	for _, src := range []string{
		"//lint:ignore SA1000 reason",
		"//lint:ignore  SA1000 reason",
		"//lint:ignore\tSA1000\treason",
	} {
		command, args := parseDirective(src)
		if command != "ignore" || !reflect.DeepEqual(args, []string{"SA1000", "reason"}) {
			t.Errorf("parseDirective(%q) = %q, %#v", src, command, args)
		}
	}
}
