package darwin_test

import (
	"testing"

	darwin "github.com/arjungandhi/darwin/pkg"
	"gopkg.in/yaml.v3"
)

func TestNode(t *testing.T) {
	// create a node object
	n := darwin.Node{
		Name:        "test",
		Description: "test node",
		Parents:     []darwin.Node{},
		Levels:      []int{1, 4, 10, 20},
		Unit:        "test",
	}

	// try to yaml marshal the node
	d, err := yaml.Marshal(n)
	if err != nil {
		t.Errorf("Error marshaling node: %s", err)
	}
	t.Logf("Marshalled node:\n%s", d)

	// try to yaml unmarshal the node
	var n2 darwin.Node
	err = yaml.Unmarshal(d, &n2)
	if err != nil {
		t.Errorf("Error unmarshaling node: %s", err)
	}
}
