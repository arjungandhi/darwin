package visualize

import (
	"io/ioutil"

	"github.com/goccy/go-graphviz"
	Z "github.com/rwxrob/bonzai/z"

	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/store"
	"github.com/rwxrob/help"
)

func init() {
	Z.Vars.SoftInit()
}

// Cmd is the root command
var Cmd = &Z.Cmd{
	Name:    "visualize",
	Summary: `renders the skill graph as a image and opens it in the default viewer`,
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
		// convert the darwin tree to a graph
		graph, err := d.ToGraph()
		if err != nil {
			return err
		}

		g := graphviz.New()

		// make a temporary image file to store the graph
		f, err := ioutil.TempFile("", "darwin-*.png")
		if err != nil {
			return err
		}
		defer f.Close()

		// render the graph to the image file
		err = g.RenderFilename(graph, graphviz.PNG, f.Name())
		if err != nil {
			return err
		}

		// open the image file in the default viewer
		return Z.Exec("xdg-open", f.Name())
	},
}
