package star

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
	Name:    "star",
	Summary: `stars a unstarred node`,
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

		n, err := search.SearchList(d.GetUnstarred())
		if err != nil {
			return err
		}

		d.Star(n.Id)
		if err != nil {
			return err
		}

		return nil
	},
}
