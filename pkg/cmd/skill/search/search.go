package search

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/google/uuid"
)

func Search(d *darwin.Darwin) (*node.Node, error) {

	// get a map of node.uuids to node.name
	nodeNameMap := make(map[string]uuid.UUID)
	nodeNames := make([]string, len(d.Nodes)+1)

	i := 0
	for u, n := range d.Nodes {
		nodeNameMap[n.Name] = u
		nodeNames[i] = n.Name
		i++
	}

	// prompt the user for the needed fields
	var answer string

	qs := &survey.Select{
		Message: "Type to search for skills by name",
		Options: nodeNames,
	}

	err := survey.AskOne(qs, &answer)
	if err != nil {
		return nil, err
	}

	// use the nodeNameMap to get the uuid of the node
	nodeUUID := nodeNameMap[answer]
	node := d.Nodes[nodeUUID]

	return node, nil
}
