// the node package provides an implemenation of a darwin node
package darwin

type Node struct {
	// Name of the node
	Name string `yaml:"name"`
	// Description of the node
	Description string `yaml:"description"`
	// Title you unlock when you achieve a node
	Title string `yaml:"title"`
	// Levels points needed to define each level of the node
	Levels []int `yaml:"levels"`
	// Unit this a type of the this is useful when displaying info
	// to a user eg.{Current Level: 100 <unit>}
	Unit string `yaml:"unit"`
	// Parents of the node
	Parents []Node `yaml:"parents"`
	// points
	Points int `yaml:"points"`
}
