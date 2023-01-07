package progress

import (
	"fmt"

	"github.com/arjungandhi/darwin/pkg/cmd/skill/search"
	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/arjungandhi/darwin/pkg/store"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"

	progressbar "github.com/tj/go-progress"
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

		// search for a skill
		node, err := search.Search(d)
		if err != nil {
			return err
		}
		// if found, print out the progress
		err = Progress(node)
		if err != nil {
			return err
		}

		return nil
	},
}

func Progress(n *node.Node) error {
	// get the NextLevelPoints
	nextLevel := n.NextLevel()
	// make a progress bar
	bar := progressbar.NewInt(n.Levels[nextLevel])

	// set the current progress
	bar.ValueInt(n.Points)

	// set the text
	bar.Text(fmt.Sprintf("%d/%d %s", n.Points, n.Levels[nextLevel], n.Unit))

	// print the progress bar
	fmt.Println(bar.String())

	return nil
}
