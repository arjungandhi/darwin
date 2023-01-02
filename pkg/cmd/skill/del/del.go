package del

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/arjungandhi/darwin/pkg/cmd/skill/search"
	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/store"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

func init() {
	Z.Vars.SoftInit()
}

// Cmd is the root command
var Cmd = &Z.Cmd{
	Name:    "del",
	Summary: `del let you delete a node from the tree`,
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Call: func(cmd *Z.Cmd, args ...string) error {
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

		// confirm that the user wants to delete the node
		qs := &survey.Confirm{
			Message: fmt.Sprintf("Are you sure you want to delete %s?", node.Name),
		}

		var answer bool
		err = survey.AskOne(qs, &answer)
		if err != nil {
			return err
		}

		if answer {
			// delete the node
			err = d.Delete(node)
		}
		return nil
	},
}
