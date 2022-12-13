// Darwin is a program that creates a skill tree for your life
// the darwin package the skill tree and the related methodology
package darwin

import (
	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/arjungandhi/darwin/pkg/store"
	"github.com/google/uuid"
)

type Darwin struct {
	// Nodes are the nodes in the skill tree
	// Nodes are stored in a map for easy access
	Nodes map[uuid.UUID]*node.Node `json:"nodes"`
	// Store is the store used to load, save, and delete nodes
	Store store.Store `json:"store"`
}

// Load creates a new darwin object and loads all the nodes from the the node store
func Load(store store.Store) (*Darwin, error) {
	n, err := store.LoadAll()
	if err != nil {
		return nil, err
	}
	d := &Darwin{
		Store: store,
		Nodes: n,
	}
	return d, nil
}
