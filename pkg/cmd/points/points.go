package points

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	Z "github.com/rwxrob/bonzai/z"

	"github.com/arjungandhi/darwin/pkg/cmd/skill/search"
	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/store"
	"github.com/rwxrob/help"
)

func init() {
	Z.Vars.SoftInit()
}

// Cmd is the root command
var Cmd = &Z.Cmd{
	Name:    "points",
	Summary: `allow you to add points to a starred skill`,
	Commands: []*Z.Cmd{
		help.Cmd,
		addCmd,
		subCmd,
	},
}

// addCmd is the add command
var addCmd = &Z.Cmd{
	Name:    "add",
	Summary: `add points to a skill`,
	Call: func(cmd *Z.Cmd, args ...string) error {
		d, s, err := setup()
		if err != nil {
			return err
		}

		n, err := search.SearchList(d.GetStarred())
		if err != nil {
			return err
		}
		points, err := askPoints("add")

		if err != nil {
			return err
		}

		n.AddPoints(points)

		err = s.Save(n)
		if err != nil {
			return err
		}

		return nil

	},
}

// subCmd is the subtract command
var subCmd = &Z.Cmd{
	Name:    "subtract",
	Summary: `subtract points from a skill`,
	Call: func(cmd *Z.Cmd, args ...string) error {

		d, s, err := setup()
		if err != nil {
			return err
		}

		n, err := search.SearchList(d.GetStarred())
		if err != nil {
			return err
		}
		points, err := askPoints("subtract")

		if err != nil {
			return err
		}

		n.AddPoints(points)

		err = s.Save(n)
		if err != nil {
			return err
		}

		return nil

	},
}

func setup() (*darwin.Darwin, store.Store, error) {
	// setup the darwin object
	s := &store.LocalStore{Path: Z.Vars.Get(".darwin.dir")}
	d, err := darwin.Load(
		s,
	)

	if err != nil {
		return nil, nil, err
	}

	return d, s, nil

}

func askPoints(word string) (int, error) {
	var answer int

	err := survey.AskOne(&survey.Input{
		Message: fmt.Sprintf("How many points do you want to %s?", word),
	}, &answer, survey.WithValidator(survey.Required))
	if err != nil {
		return 0, err
	}

	return answer, nil
}
