package edit

import (
	"github.com/AlecAivazis/survey/v2"
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
	Name:    "edit",
	Summary: `opens the text editor for you to edit node properties`,
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

		// use survey.Editor to edit the node
		// serialize the node to a yaml string
		yamlNode, err := yaml.Marshal(node)
		if err != nil {
			return err
		}
		// then use survery.Editor to open the yaml in the editor
		prompt := &survey.Editor{
			Message:       "Edit the node",
			Default:       string(yamlNode),
			AppendDefault: true,
			FileName:      "*.yaml",
		}

		var answer string

		survey.AskOne(prompt, &answer)

		// then deserialize the yaml back into a node
		err = yaml.Unmarshal([]byte(answer), &node)
		if err != nil {
			return err
		}

		// use the store to save the updated node
		err = dStore.Save(node)

		return nil
	},
}
