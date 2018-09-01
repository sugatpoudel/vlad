package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	isPositive = &Node{
		Validate: func(num int) bool { return num > -1 },
		Msg:      "Number is not positive",
	}

	isEven = &Node{
		Validate: func(num int) bool { return num%2 == 0 },
		Msg:      "Number is not even",
	}

	isMultipleOfThree = &Node{
		Validate: func(num int) bool { return num%3 == 0 },
		Msg:      "Number is not a multiple of 3",
	}
)

func TestNewGraph(t *testing.T) {
	graph := NewGraph()
	assert.NotNil(t, graph, "Did not create graph properly")
	assert.NotNil(t, graph.nodes, "Did not create nodes properly")
	assert.NotNil(t, graph.edges, "Did not create edges properly")
}

func TestAddNode(t *testing.T) {
	graph := NewGraph()
	added := graph.AddNode(isPositive)

	assert.True(t, added, "Node was not added to the graph")
	assert.Equal(t, 1, len(graph.GetNodes()), "Node was not added to the nodes list")
}

func TestPutEdge(t *testing.T) {
	graph := NewGraph()
	added := graph.PutEdge(isPositive, isEven)

	assert.True(t, added, "Edge was not added to the graph")
	assert.Equal(t, 2, len(graph.GetNodes()), "Nodes were not added to the graph")
	assert.Equal(t, 1, len(graph.GetDependents(isPositive)), "Wrong number of dependents")
}

func TestPutTwoEdges(t *testing.T) {
	graph := NewGraph()
	graph.PutEdge(isPositive, isEven)
	added := graph.PutEdge(isPositive, isMultipleOfThree)

	assert.True(t, added, "Second edge was not added to the graph")
	assert.Equal(t, 3, len(graph.GetNodes()), "All nodes were not added to the graph")
	assert.Equal(t, 2, len(graph.GetDependents(isPositive)), "Wrong number of dependents")
}

func TestValidateWithOneLevel(t *testing.T) {
	graph := NewGraph()
	graph.PutEdge(isPositive, isEven)

	errors := graph.Validate(-1)
	assert.NotNil(t, errors, "Errors were null")
	assert.Equal(t, 1, len(errors), "Validate returned no errors")
	assert.Equal(t, "Number is not positive", errors[0], "Incorrect Validation")
}

func TestValidateWithTwoLevels(t *testing.T) {
	graph := NewGraph()
	graph.PutEdge(isPositive, isEven)
	graph.PutEdge(isPositive, isMultipleOfThree)

	errors := graph.Validate(7)
	assert.NotNil(t, errors, "Errors were null")
	assert.Equal(t, 2, len(errors), "Validate returned no errors")
}
