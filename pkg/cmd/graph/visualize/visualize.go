package visualize

import (
	"encoding/json"
	"os"
	"text/template"

	"golang.org/x/exp/maps"

	Z "github.com/rwxrob/bonzai/z"

	"github.com/arjungandhi/darwin/pkg/darwin"
	"github.com/arjungandhi/darwin/pkg/store"
	"github.com/rwxrob/help"
)

func init() {
	Z.Vars.SoftInit()
}

var visTmpl = template.Must(template.New("vis").Parse(`<html>
  <head>
    <style>
      body {
        width: 100%;
        height: 100%;
        margin: 0;
      }
    </style>
    <meta charset="utf-8" />
    <script src="https://d3js.org/d3.v7.min.js"></script>
    <script
      src="https://cdn.jsdelivr.net/gh/arjungandhi/d3-darwin/d3.darwin.js"
      charset="utf-8"
    ></script>
  </head>
  <body></body>

  <script type="text/javascript">
    var skills = {{.}};
    // make a canvas element
    var canvas = document.createElement("canvas");
    canvas.style.width = "100%";
    canvas.style.height = "100%";
    document.body.appendChild(canvas);

    // create the darwin object
    var darwin = d3
      .darwin()
      .skills(skills)
      .height(document.body.offsetHeight)
      .width(document.body.offsetWidth)
      .canvas(canvas);

    // setup a listener to listen for resize events and update height and width when that happens
    window.addEventListener("resize", function () {
      darwin
        .height(document.body.offsetHeight)
        .width(document.body.offsetWidth);
      darwin();
    });

    // run the darwin object
    darwin();
  </script>
</html>`))

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

		// convert to json
		json, err := json.Marshal(maps.Values(d.Nodes))
		if err != nil {
			return err
		}

		// setup a temp file and write the html to it
		f, err := os.CreateTemp("", "darwin-visualize-*.html")
		if err != nil {
			return err
		}
		defer f.Close()

		// write the html to the temp file
		err = visTmpl.Execute(f, string(json))
		if err != nil {
			return err
		}

		// user the default browser to open the temp file
		return Z.Exec("xdg-open", f.Name())
		return nil
	},
}
