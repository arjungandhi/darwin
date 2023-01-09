package skill

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"

	"github.com/arjungandhi/darwin/pkg/cmd/skill/add"
	"github.com/arjungandhi/darwin/pkg/cmd/skill/del"
	"github.com/arjungandhi/darwin/pkg/cmd/skill/edit"
	"github.com/arjungandhi/darwin/pkg/cmd/skill/info"
	"github.com/arjungandhi/darwin/pkg/cmd/skill/points"
	"github.com/arjungandhi/darwin/pkg/cmd/skill/progress"
	"github.com/arjungandhi/darwin/pkg/cmd/skill/star"
	"github.com/arjungandhi/darwin/pkg/cmd/skill/unstar"
)

func init() {
	Z.Vars.SoftInit()
}

// Cmd is the root command for the skill functions
var Cmd = &Z.Cmd{
	Name:    "skill",
	Summary: `The node cmd provides interfaces for interacting with darwin trees`,
	Commands: []*Z.Cmd{
		help.Cmd,
		add.Cmd,
		del.Cmd,
		info.Cmd,
		edit.Cmd,
		progress.Cmd,
		star.Cmd,
		unstar.Cmd,
		points.Cmd,
	},
}
