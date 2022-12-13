// the node_store creates and implements various methods to store
// and retrieve nodes.
// it's setup to be an interface so that it can be easily swapped out with other
// implementations as darwin grows
package store

import (
	"os"
	"path/filepath"

	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
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

type LocalStore struct {
	// Path is the path to the local store
	Path string `json:"path"`
}

// Save saves a node to the store
func (s *LocalStore) Save(n *node.Node) error {
	path := filepath.Join(s.Path, n.Id.String()+".yaml")
	// marshal the node to yaml
	d, err := yaml.Marshal(n)
	if err != nil {
		return err
	}
	// write the yaml to the file
	err = os.WriteFile(path, d, 0666)
	if err != nil {
		return err
	}
	return nil
}

// Load loads a node from the store
func (s *LocalStore) Load(id uuid.UUID) (*node.Node, error) {
	path := filepath.Join(s.Path, id.String()+".yaml")
	// read the file
	d, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// unmarshal the yaml to a node
	var n node.Node
	err = yaml.Unmarshal(d, &n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// LoadAll searches the Path for all nodes that follow the
// *.yaml format and them loads them into a slice of nodes
func (s *LocalStore) LoadAll() (map[uuid.UUID]*node.Node, error) {
	// search the path for all files that follow the *.yaml format
	matches, err := filepath.Glob(filepath.Join(s.Path, "*.yaml"))
	if err != nil {
		return nil, err
	}
	// map of nodes
	nodes := make(map[uuid.UUID]*node.Node)
	// loop through all the matches
	for _, match := range matches {
		// read the file
		d, err := os.ReadFile(match)
		if err != nil {
			return nil, err
		}
		// unmarshal the yaml to a node
		var n node.Node
		err = yaml.Unmarshal(d, &n)
		if err != nil {
			return nil, err
		}
		// add the node to the map
		nodes[n.Id] = &n
	}

	return nodes, nil
}

// Delete deletes a node from the store
func (s *LocalStore) Delete(n *node.Node) error {
	// create the path to the file
	path := filepath.Join(s.Path, n.Id.String()+".yaml")
	// delete the file
	return os.Remove(path)
}
