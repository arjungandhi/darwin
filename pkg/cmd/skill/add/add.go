package add

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/arjungandhi/darwin/pkg/store"
	"github.com/google/uuid"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"gopkg.in/yaml.v3"
)

func init() {
	Z.Vars.SoftInit()
}

// Cmd is the root command
var Cmd = &Z.Cmd{
	Name:    "add",
	Summary: `The node cmd provides interfaces for interacting with darwin trees`,
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Call: func(cmd *Z.Cmd, args ...string) error {
		return AddSkill()
	},
}

func AddSkill() error {
	// setup the darwin object
	dStore := &store.LocalStore{Path: Z.Vars.Get(".darwin.dir")}
	d, err := darwin.Load(
		dStore,
	)

	if err != nil {
		return err
	}

	// get a map of node.uuids to node.name
	nodeNameMap := make(map[string]uuid.UUID)
	nodeNames := make([]string, len(d.Nodes)+1)
	//nodeNames[0] = "None"
	//nodeNameMap["None"] = uuid.Nil

	i := 0
	for u, n := range d.Nodes {
		nodeNameMap[n.Name] = u
		nodeNames[i] = n.Name
		i++
	}

	// prompt the user for the needed fields
	answers := struct {
		Name   string
		Parent []string
		Edit   bool
	}{}

	qs := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What is the name of the skill?",
			},
			Validate: survey.Required,
		},
		{
			Name: "parent",
			Prompt: &survey.MultiSelect{
				Message: "What is the parent of the skill?",
				Options: nodeNames,
			},
		},
		{
			Name: "edit",
			Prompt: &survey.Confirm{
				Message: "Do you want to edit the skill?",
			},
		},
	}

	err = survey.Ask(qs, &answers)
	if err != nil {
		return err
	}

	// use the nodeNameMap to get the uuid of the parent
	parentUUIDs := []uuid.UUID{}
	for p := range answers.Parent {
		u := nodeNameMap[answers.Parent[p]]
		if u == uuid.Nil {
			continue
		}
		parentUUIDs = append(parentUUIDs, u)
	}
	// create the node
	n := node.New(answers.Name, parentUUIDs)

	// if edit is true, open the node in the editor
	if answers.Edit {
		// serialize the node to a yaml string
		yamlNode, err := yaml.Marshal(n)
		if err != nil {
			return err
		}
		// then use survery.Editor to open the yaml in the editor
		prompt := &survey.Editor{
			Message:       "Edit the node, you can edit all node properties",
			Default:       string(yamlNode),
			AppendDefault: true,
			FileName:      "*.yaml",
		}

		var answer string

		survey.AskOne(prompt, &answer)

		// then deserialize the yaml back into a node
		err = yaml.Unmarshal([]byte(answer), &n)
		if err != nil {
			return err
		}
	}

	// add the node to the darwin object
	err = d.Add(n)
	if err != nil {
		return err
	}

	return nil

}
