// Darwin is a program that creates a skill tree for your life
// the darwin package the skill tree and the related methodology
package darwin

import (
	"fmt"

	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/arjungandhi/darwin/pkg/store"
	"github.com/google/uuid"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type Darwin struct {
	// Nodes are the nodes in the skill tree
	// Nodes are stored in a map for easy access
	Nodes map[uuid.UUID]*node.Node
	// Store is the store used to load, save, and delete nodes
	Store store.Store
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

// Adds a node to the darwin tree
func (d *Darwin) Add(n *node.Node) error {
	d.Nodes[n.Id] = n
	// save the node in the store
	err := d.Store.Save(n)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a node from the darwin tree
func (d *Darwin) Delete(n *node.Node) error {
	delete(d.Nodes, n.Id)
	d.Store.Delete(n)
	return nil
}

// Star adds a node to the starred list
func (d *Darwin) Star(u uuid.UUID) error {
	n, ok := d.Nodes[u]
	if !ok {
		return fmt.Errorf("node with uuid %s not found", u)
	}
	n.Starred = true
	d.Store.Save(n)
	return nil
}

// Unstar removes a node from the starred list
func (d *Darwin) Unstar(u uuid.UUID) error {
	n, ok := d.Nodes[u]
	if !ok {
		return fmt.Errorf("node with uuid %s not found", u)
	}
	n.Starred = false
	d.Store.Save(n)
	return nil
}

// GetStarred returns a list of starred nodes
func (d *Darwin) GetStarred() []*node.Node {
	var starred []*node.Node
	for _, n := range d.Nodes {
		if n.Starred {
			starred = append(starred, n)
		}
	}
	return starred
}

func (d *Darwin) ToGraph() (*cgraph.Graph, error) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return nil, err
	}

	nodes := make(map[uuid.UUID]*cgraph.Node)
	// add nodes to the graph
	for _, n := range d.Nodes {
		node, err := graph.CreateNode(n.Id.String())
		nodes[n.Id] = node
		node.SetLabel(n.Name)
		if err != nil {
			return nil, err
		}
	}
	// add edges to the graph
	for _, n := range d.Nodes {
		for _, p := range n.Parents {
			_, err := graph.CreateEdge("", nodes[p], nodes[n.Id])
			if err != nil {
				return nil, err
			}
		}
	}

	return graph, nil

}
