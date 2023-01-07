// the node package provides an implemenation of a darwin node
package node

import (
	"time"

	"github.com/google/uuid"
)

type Node struct {
	// Id is a unique identifier for the node
	Id uuid.UUID `json:"id" yaml:"id"`
	// Name of the node
	Name string `json:"name" yaml:"name"`
	// Description of the node
	Description string `json:"description" yaml:"description"`
	// Title you unlock when you achieve a node
	Title string `json:"title" yaml:"title"`
	// Levels points needed to define each level of the node
	// all levels start at 0
	Levels []int `json:"levels" yaml:"levels"`
	// Unit this a type of the this is useful when displaying info
	// to a user eg.{Current Level: 100 <unit>}
	Unit string `json:"unit" yaml:"unit"`
	// Parents of the node
	Parents     []uuid.UUID `json:"parents"	yaml:"parents"`
	ParentNodes []*Node     `json:"-" yaml:"-"`
	// points are the current points the user has achieved
	Points int `json:"points" yaml:"points"`
	// LastAchieved is the unix timestamp of the last time a level was achieved
	LastAchieved int64 `json:"last_achieved" yaml:"last_achieved"`
	// Starred is a boolean value that represents if the node is starred
	Starred bool `json:"starred" yaml:"starred"`
}

// NewNode creates a new node
// with sane defaults
func New(name string, parents []uuid.UUID) *Node {
	return &Node{
		Id:           uuid.New(),
		Name:         name,
		Description:  "",
		Title:        name + "er",
		Levels:       []int{0, 1, 5, 10, 20, 50},
		Unit:         "points",
		Parents:      parents,
		ParentNodes:  []*Node{},
		Points:       0,
		LastAchieved: time.Now().Unix(),
		Starred:      false,
	}
}

// Progress returns the a value between 0 and 1 representing the
// the amount of points achieved towards the next level
func (n *Node) Progress() float32 {
	currentLevel := n.Level()
	nexLevel := n.NextLevel()
	if currentLevel == nexLevel {
		return 1
	}
	nextLevelPoints := n.Levels[nexLevel]
	return float32(n.Points) / float32(nextLevelPoints)
}

// Level returns the current level of the node
func (n *Node) Level() int {
	currentLevel := 0
	for i, level := range n.Levels {
		if n.Points < level {
			break
		}
		currentLevel = i
	}
	return currentLevel
}

// NextLevel returns the nextLevel the node could achieve
// or the max level if the node is already at the max level
func (n *Node) NextLevel() int {
	currentLevel := n.Level()
	if currentLevel == len(n.Levels)-1 {
		return currentLevel
	}
	return currentLevel + 1
}

// AddPoints adds points to the node and updates the last achieved time
func (n *Node) AddPoints(points int) {
	oldLevel := n.Level()
	n.Points += points
	if n.Level() != oldLevel {
		n.LastAchieved = time.Now().UnixNano()
	}

	// recursively update the parents
	for _, parent := range n.ParentNodes {
		parent.AddPoints(points)
	}
}
