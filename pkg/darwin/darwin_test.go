package darwin_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/arjungandhi/darwin/pkg/store"
	"github.com/google/uuid"
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

	// test adding a node
	node := &node.Node{
		Id:   uuid.New(),
		Name: "test",
	}

	t.Run("Add", func(t *testing.T) {
		err := darwinTree.Add(node)
		if err != nil {
			t.Errorf("Error adding node: %s", err)
		}
		// check if the node was added
		_, ok := darwinTree.Nodes[node.Id]
		if !ok {
			t.Errorf("Node not added")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := darwinTree.Delete(node)
		if err != nil {
			t.Errorf("Error deleting node: %s", err)
		}
		// check if the node was deleted
		_, ok := darwinTree.Nodes[node.Id]
		if ok {
			t.Errorf("Node not deleted")
		}
	})
}
