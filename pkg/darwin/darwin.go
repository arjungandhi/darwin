// Darwin is a program that creates a skill tree for your life
// the darwin package the skill tree and the related methodology
package darwin

import (
	"fmt"

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
	// Starred Nodes are a list of UUIDs for "starred" nodes
	// starred nodes are meant to be easily accesible for the node owner
	Starred []uuid.UUID
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

func (d *Darwin) Add(n *node.Node) error {
	d.Nodes[n.Id] = n
	// save the node in the store
	err := d.Store.Save(n)
	if err != nil {
		return err
	}

	return nil
}

func (d *Darwin) Delete(n *node.Node) error {
	delete(d.Nodes, n.Id)
	d.Store.Delete(n)
	return nil
}

func (d *Darwin) Star(u uuid.UUID) error {
	n, ok := d.Nodes[u]
	if !ok {
		return fmt.Errorf("node with uuid %s not found", u)
	}
	d.Starred = append(d.Starred, n.Id)
	return nil
}

func (d *Darwin) UnStar(u uuid.UUID) error {
	// search for the node in the starred list
	for i, v := range d.Starred {
		if v == u {
			// remove the node from the starred list
			d.Starred = append(d.Starred[:i], d.Starred[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("node with uuid %s not found", u)
}
