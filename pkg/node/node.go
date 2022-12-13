// the node package provides an implemenation of a darwin node
package node

import (
	"github.com/google/uuid"
)

type Node struct {
	// Id is a unique identifier for the node
	Id uuid.UUID `json:"id"`
	// Name of the node
	Name string `json:"name"`
	// Description of the node
	Description string `json:"description"`
	// Title you unlock when you achieve a node
	Title string `json:"title"`
	// Levels points needed to define each level of the node
	Levels []int `json:"levels"`
	// Unit this a type of the this is useful when displaying info
	// to a user eg.{Current Level: 100 <unit>}
	Unit string `json:"unit"`
	// Parents of the node
	Parents []uuid.UUID `json:"parents"`
	// points are the current points the user has achieved
	Points int `json:"points"`
}
