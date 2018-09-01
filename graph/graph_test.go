package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	node1 = &Node{
		validate: func(a int) bool { return false },
		msg:      "Failed validation 1",
	}
	node2 = &Node{
		validate: func(a int) bool { return true },
		msg:      "Failed validation 2",
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
	added := graph.AddNode(node1)

	assert.True(t, added, "Node was not added to the graph")
	assert.Equal(t, 1, len(graph.GetNodes()), "Node was not added to the nodes list")
}

func TestPutEdge(t *testing.T) {
	graph := NewGraph()
	added := graph.PutEdge(node1, node2)

	assert.True(t, added, "Edge was not added to the graph")
	assert.Equal(t, 2, len(graph.GetNodes()), "Nodes were not added to the graph")
}
