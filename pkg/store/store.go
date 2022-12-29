// the node_store creates and implements various methods to store
// and retrieve nodes.
// it's setup to be an interface so that it can be easily swapped out with other
// implementations as darwin grows
package store

import (
	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/google/uuid"
)

// The NodeStore interface defines the methods that a node object
// will use to store, edit and retrieve nodes.
type Store interface {
	// Save saves a node to its store
	Save(*node.Node) error
	// Load loads a node from its store
	Load(uuid.UUID) (*node.Node, error)
	// LoadAll loads all the nodes from its store
	LoadAll() (map[uuid.UUID]*node.Node, error)
	// Delete deletes a node from its store
	Delete(*node.Node) error
}
