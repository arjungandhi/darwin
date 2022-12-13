package store_test

import (
	"os"
	"testing"

	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/arjungandhi/darwin/pkg/store"
	"github.com/google/uuid"
)

var testNodes = []*node.Node{
	{
		Id:          uuid.New(),
		Name:        "test",
		Description: "test node",
		Parents: []uuid.UUID{
			uuid.New(),
			uuid.New(),
		},
		Levels: []int{1, 4, 10, 20},
		Unit:   "test",
	},
	{
		Id:          uuid.New(),
		Name:        "test2",
		Description: "test node 2",
		Parents:     []uuid.UUID{},
		Levels:      []int{1, 4, 10, 20},
		Unit:        "test",
	},
}

func TestLocalStore(t *testing.T) {
	// create a temporary directory
	dir, err := os.MkdirTemp("", "*")
	if err != nil {
		t.Errorf("Error creating temporary directory: %s", err)
	}
	defer os.RemoveAll(dir)
	// create a local store object
	store := store.LocalStore{
		Path: dir,
	}
	// iterate over the test nodes
	for _, n := range testNodes {
		// save the node to the store
		err = store.Save(n)
		if err != nil {
			t.Errorf("Error saving node: %s", err)
		}
	}

	// load the node from the store
	for _, n := range testNodes {
		_, err := store.Load(n.Id)
		if err != nil {
			t.Errorf("Error loading node: %s", err)
		}
	}

	// load all the nodes from the store
	_, err = store.LoadAll()
	if err != nil {
		t.Errorf("Error loading all nodes: %s", err)
	}

	// delete the node from the store
	for _, n := range testNodes {
		err = store.Delete(n)
		if err != nil {
			t.Errorf("Error deleting node: %s", err)
		}
	}
}
