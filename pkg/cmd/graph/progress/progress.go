package progress

import (
	sprogress "github.com/arjungandhi/darwin/pkg/cmd/skill/progress"
	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/arjungandhi/darwin/pkg/store"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

func init() {
	Z.Vars.SoftInit()
}

// Cmd is the root command
var Cmd = &Z.Cmd{
	Name:    "progress",
	Summary: `Prints out a skills progress`,
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Call: func(cmd *Z.Cmd, args ...string) error {
		// setup the darwin tree
		dStore := &store.LocalStore{Path: Z.Vars.Get(".darwin.dir")}
		d, err := darwin.Load(
			dStore,
		)
		if err != nil {
			return err
		}

		// get starred nodes
		starredNodes := []*node.Node{}
		for _, n := range d.Nodes {
			if n.Starred {
				starredNodes = append(starredNodes, n)
			}
		}
		for _, n := range starredNodes {
			// if found, print out the progress
			err = sprogress.Progress(n)
			if err != nil {
				return err
			}
		}

		return nil
	},
}
