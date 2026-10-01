package gocommand

import (
	"path/filepath"
	"testing"
)

func TestWorkspaceVendorDir(t *testing.T) {
	for _, goWork := range []string{"", "off"} {
		if dir, ok := workspaceVendorDir(goWork); ok {
			t.Errorf("workspaceVendorDir(%q) = %q, true; want false", goWork, dir)
		}
	}
	if dir, ok := workspaceVendorDir(filepath.Join("work", "go.work")); !ok || dir != filepath.Join("work", "vendor") {
		t.Errorf("workspaceVendorDir(work/go.work) = %q, %t; want work/vendor, true", dir, ok)
	}
}
