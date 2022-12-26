package darwin_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/store"
	"gopkg.in/yaml.v3"
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

	// Run Test to see if we can marshal and unmarshal the darwin object
	t.Run("TestMarshal", func(t *testing.T) {
		// print out the yaml representation of the darwin object
		d, err := yaml.Marshal(darwinTree)
		if err != nil {
			t.Errorf("Error marshaling darwin: %s", err)
		}
		t.Logf("Marshalled darwin:\n%s", d)

		// unmarshal the yaml representation of the darwin object
		var darwinTree2 darwin.Darwin
		err = yaml.Unmarshal(d, &darwinTree2)
		if err != nil {
			t.Errorf("Error unmarshaling darwin: %s", err)
		}
	})

	// Test addi
}
