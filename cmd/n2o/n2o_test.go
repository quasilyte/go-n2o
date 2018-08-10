package main

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestOutput(t *testing.T) {
	have := execCommand(t, "go", "run", filepath.Join("testdata", "main.n2o.go"))
	want := execCommand(t, "go", "run", filepath.Join("testdata", "main.go"))
	if have != want {
		t.Errorf("output mismatches:\nhave: %s\nwant: %s", have, want)
	}
}

// execCommand runs exec.Command and returns combined output as a string.
// Execution errors result in t.Fatalf call.
func execCommand(t *testing.T, name string, args ...string) string {
	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		t.Fatalf("exec %s failed (%v): %s", name, err, out)
	}
	return string(out)
}
