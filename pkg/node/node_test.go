package node_test

import (
	"testing"
	"time"

	"github.com/arjungandhi/darwin/pkg/node"
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

func TestNode(t *testing.T) {
	// create a node object to use for testing
	n := node.Node{
		Id:          uuid.New(),
		Name:        "test",
		Description: "test node",
		Parents: []uuid.UUID{
			uuid.New(),
		},
		Levels:       []int{0, 1, 4, 10, 20},
		Unit:         "test",
		LastAchieved: time.Now().Unix(),
	}
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
		{1, 0},
		{2, 1.0 / 3.0},
		{3, 2.0 / 3.0},
		{4, 0},
		{5, 1.0 / 6.0},
		{8, 4.0 / 6.0},
		{9, 5.0 / 6.0},
		{10, 0},
		{11, 1.0 / 10.0},
		{12, 2.0 / 10.0},
		{13, 3.0 / 10.0},
		{14, 4.0 / 10.0},
		{15, 5.0 / 10.0},
		{16, 6.0 / 10.0},
		{17, 7.0 / 10.0},
		{18, 8.0 / 10.0},
		{19, 9.0 / 10.0},
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

}
