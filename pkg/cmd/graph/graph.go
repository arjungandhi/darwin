package graph

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"

	"github.com/arjungandhi/darwin/pkg/cmd/graph/progress"
	"github.com/arjungandhi/darwin/pkg/cmd/graph/visualize"
)

func init() {
	Z.Vars.SoftInit()
}

// Cmd is the root command for the skill functions
var Cmd = &Z.Cmd{
	Name:    "graph",
	Summary: `graph provides methods for interacting with the skill graph`,
	Commands: []*Z.Cmd{
		help.Cmd,
		visualize.Cmd,
		progress.Cmd,
	},
}
