package info

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"

	"github.com/arjungandhi/darwin/pkg/cmd/skill/search"
	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/store"
	"github.com/rwxrob/help"
	"gopkg.in/yaml.v3"
)

func init() {
	Z.Vars.SoftInit()
}

// Cmd is the root command
var Cmd = &Z.Cmd{
	Name:    "info",
	Summary: `prints out the yaml representation of a node`,
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Call: func(cmd *Z.Cmd, args ...string) error {
		// setup the darwin object
		dStore := &store.LocalStore{Path: Z.Vars.Get(".darwin.dir")}
		d, err := darwin.Load(
			dStore,
		)

		if err != nil {
			return err
		}
		node, err := search.Search(d)
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
