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

	node := &node.Node{
		Id:      uuid.New(),
		Name:    "test",
		Starred: false,
	}

	// test adding a node
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

	// test getting the list of starred nodes
	t.Run("GetStarred", func(t *testing.T) {
		starredNodes := darwinTree.GetStarred()
		if len(starredNodes) != 1 {
			t.Errorf("Incorrect number of starred nodes")
		}
	})

	// test staring a node
	t.Run("Star", func(t *testing.T) {
		err := darwinTree.Star(node.Id)
		if err != nil {
			t.Errorf("Error starring node: %s", err)
		}
		// check if the node was starred
		n, ok := darwinTree.Nodes[node.Id]
		if !ok {
			t.Errorf("Node not found")
		}
		if !n.Starred {
			t.Errorf("Node not starred")
		}
	})
	// test unstarring a node
	t.Run("Unstar", func(t *testing.T) {
		err := darwinTree.Unstar(node.Id)
		if err != nil {
			t.Errorf("Error unstarring node: %s", err)
		}
		// check if the node was unstarred
		n, ok := darwinTree.Nodes[node.Id]
		if !ok {
			t.Errorf("Node not found")
		}
		if n.Starred {
			t.Errorf("Node not unstarred")
		}
	})

	// test deleting a node
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
