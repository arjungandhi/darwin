package darwin_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/store"
)

func TestDarwin(t *testing.T) {
	// Run Load to create a new darwin object
	_, filename, _, _ := runtime.Caller(0)
	testDir := filepath.Join(filepath.Dir(filename), "testdata")

	darwinTree, err := darwin.Load(
		&store.LocalStore{
			Path: testDir,
		},
	)
	if err != nil {
		t.Errorf("Error loading darwin: %s", err)
	}

}
