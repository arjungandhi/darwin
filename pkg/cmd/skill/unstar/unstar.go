package unstar

import (
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
	Name:    "unstar",
	Summary: `unstars a starred node`,
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

		n, err := search.SearchList(d.GetStarred())
		if err != nil {
			return err
		}

		d.Unstar(n.Id)
		err = dStore.Save(n)
		if err != nil {
			return err
		}

		return nil
	},
}
