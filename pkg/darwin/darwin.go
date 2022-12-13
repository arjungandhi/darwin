// Darwin is a program that creates a skill tree for your life
// the darwin package the skill tree and the related methodology
package darwin

import (
	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/arjungandhi/darwin/pkg/store"
)

type Darwin struct {
	// Nodes are the nodes in the skill tree
	nodes []*node.Node
	// Store is the store used to load, save, and delete nodes
	store store.Store
}

// Load creates a new darwin object and loads all the nodes from the the node store
func Load(store store.Store) (*Darwin, error) {
	return nil, nil
}
