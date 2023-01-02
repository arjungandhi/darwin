package search

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/arjungandhi/darwin/pkg/store"
	"github.com/google/uuid"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"gopkg.in/yaml.v3"
)

func init() {
	Z.Vars.SoftInit()
}

// Cmd is the root command
var Cmd = &Z.Cmd{
	Name:    "search",
	Summary: `searches for a skill in the tree, and then prints out its file`,
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Call: func(cmd *Z.Cmd, args ...string) error {
		node, err := Search()
		if err != nil {
			return err
		}

		// print out the node&
		data, err := yaml.Marshal(node)
		// remove the \n at the end of the string
		data = data[:len(data)-1]
		if err != nil {
			return err
		}
		fmt.Println(string(data))

		return nil
	},
}

func Search() (*node.Node, error) {
	// setup the darwin object
	dStore := &store.LocalStore{Path: Z.Vars.Get(".darwin.dir")}
	d, err := darwin.Load(
		dStore,
	)

	if err != nil {
		return nil, err
	}

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

	err = survey.AskOne(qs, &answer)
	if err != nil {
		return nil, err
	}

	// use the nodeNameMap to get the uuid of the node
	nodeUUID := nodeNameMap[answer]
	node := d.Nodes[nodeUUID]

	return node, nil
}
