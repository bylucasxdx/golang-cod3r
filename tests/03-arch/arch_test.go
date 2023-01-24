package arch

import (
	"runtime"
	"testing"
)

func TestDependencies(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		t.Skip("Arch AMD64, skiping test...")
	}

	t.Fail()
}
