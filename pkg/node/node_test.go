package node_test

import (
	"testing"

	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

func TestNode(t *testing.T) {
	// create a node object to use for testing
	n := node.New("test", []uuid.UUID{uuid.New()})
	n.Levels = []int{0, 1, 4, 10, 20}
	n1 := node.New("test1", []uuid.UUID{n.Id})
	n1.ParentNodes = []*node.Node{n}
	// test that the node can marshal to and from yaml
	t.Run("TestMarshal", func(t *testing.T) {

		// try to yaml marshal the node
		d, err := yaml.Marshal(n)
		if err != nil {
			t.Errorf("Error marshaling node: %s", err)
		}
		t.Logf("Marshalled node:\n%s", d)

		// try to yaml unmarshal the node
		var n2 node.Node
		err = yaml.Unmarshal(d, &n2)
		if err != nil {
			t.Errorf("Error unmarshaling node: %s", err)
		}
	})

	// test LevelFunc
	levelTests := []struct {
		points int
		level  int
	}{
		{0, 0},
		{1, 1},
		{3, 1},
		{4, 2},
		{9, 2},
		{10, 3},
		{19, 3},
		{20, 4},
		{21, 4},
	}

	for _, test := range levelTests {
		t.Run("TestCurrentLevel", func(t *testing.T) {
			n.Points = test.points
			if n.Level() != test.level {
				t.Errorf("Points{%d}, Expected level %d but got %d", test.points, test.level, n.Level())
			}
		})
	}

	// test the progress function
	progressTests := []struct {
		points   int
		progress float32
	}{
		{0, 0},
		{1, 1.0 / 4.0},
		{2, 2.0 / 4.0},
		{3, 3.0 / 4.0},
		{4, 4.0 / 10.0},
		{5, 5.0 / 10.0},
		{8, 8.0 / 10.0},
		{9, 9.0 / 10.0},
		{10, 10.0 / 20.0},
		{11, 11.0 / 20.0},
		{19, 19.0 / 20.0},
		{20, 1},
		{21, 1},
		{100, 1},
	}

	t.Run("TestProgress", func(t *testing.T) {
		for _, test := range progressTests {
			n.Points = test.points
			p := n.Progress()
			if p != test.progress {
				t.Errorf("Progress(%d) = %f, want %f", test.points, p, test.progress)
			}

		}
	})

	// test the add points function
	t.Run("TestAddPoints", func(t *testing.T) {
		n.Points = 0
		for i := 0; i < 25; i++ {
			oldTime := n.LastAchieved
			n.AddPoints(1)
			//check if the current point value is a level
			if n.Points == n.Levels[n.Level()] {
				// check if the last achieved time has changed
				if n.LastAchieved == oldTime {
					t.Errorf("AddPoints(%d) Last achieved time did not change", n.Points)
				}
			}
		}
	})

	// check that adding points also updates the parent nodes
	t.Run("TestAddPointsParent", func(t *testing.T) {
		n.Points = 0
		n1.Points = 0
		n1.AddPoints(1)
		if n.Points != 1 {
			t.Errorf("AddPoints(%d) did not update child node", n.Points)
		}
	})

}
