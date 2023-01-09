package cmd

import (
	"fmt"
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"

	"github.com/arjungandhi/darwin/pkg/cmd/graph"
	"github.com/arjungandhi/darwin/pkg/cmd/points"
	"github.com/arjungandhi/darwin/pkg/cmd/progress"
	"github.com/arjungandhi/darwin/pkg/cmd/skill"
	"github.com/arjungandhi/darwin/pkg/cmd/visualize"
)

// init runs immediately after the package is loaded.
// it sets up a few things that are needed for the command to run nicely
func init() {
	Z.Vars.SoftInit()

	// set darwinDir
	darwindir, exists := os.LookupEnv("DARWINDIR")
	if !exists {
		fmt.Println("DARWINDIR environment variable not set")
		os.Exit(1)
	}

	// set the zetdir var
	Z.Vars.Set(".darwin.dir", darwindir)
}

// RootCmd is the command that is run when the darwin command is called
var RootCmd = &Z.Cmd{
	Name:    "darwin",
	Summary: `The darwin command is a command line interface to a darwin skill tree`,
	Commands: []*Z.Cmd{
		help.Cmd,
		vars.Cmd,
		skill.Cmd,
		graph.Cmd,
		points.Cmd,
		visualize.Cmd,
		progress.Cmd,
	},
}
