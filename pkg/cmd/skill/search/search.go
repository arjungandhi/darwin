package search

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/node"
	"golang.org/x/exp/maps"
)

func Search(d *darwin.Darwin) (*node.Node, error) {
	nodes := maps.Values(d.Nodes)
	return SearchList(nodes)
}

func SearchList(nodeList []*node.Node) (*node.Node, error) {

	// get a map of node.uuids to node.name
	nodeNameMap := make(map[string]*node.Node)
	nodeNames := make([]string, len(nodeList)+1)

	for i, n := range nodeList {
		nodeNameMap[n.Name] = n
		nodeNames[i] = n.Name
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
	node := nodeNameMap[answer]

	return node, nil
}
