package graph

import (
	"github.com/golang-collections/collections/stack"
)

// Node represents a validation node
type Node struct {
	Validate func(int) bool
	Msg      string
}

// edges represents the adjacency list of graph
type edges map[*Node][]*Node

// Graph represents the validation dependencies
type Graph struct {
	nodes []*Node
	edges edges
}

// NewGraph constructs a graph
func NewGraph() *Graph {
	nodes := make([]*Node, 0)
	edges := make(map[*Node][]*Node)
	return &Graph{
		nodes: nodes,
		edges: edges,
	}
}

// GetNodes returns all nodes of the graph in insertion order
func (g *Graph) GetNodes() []*Node {
	return g.nodes
}

// AddNode adds the given node to the graph
func (g *Graph) AddNode(node *Node) bool {
	if _, ok := g.edges[node]; ok {
		return false
	}
	g.nodes = append(g.nodes, node)
	g.edges[node] = make([]*Node, 0)
	return true
}

// PutEdge adds an edge from n1 to n2
func (g *Graph) PutEdge(n1 *Node, n2 *Node) bool {
	g.AddNode(n1)
	g.AddNode(n2)
	found := false
	for _, n := range g.edges[n1] {
		if n == n2 {
			found = true
		}
	}
	if !found {
		g.edges[n1] = append(g.edges[n1], n2)
		return true
	}
	return false
}

// GetDependents retrives all dependents of the given node
func (g *Graph) GetDependents(node *Node) []*Node {
	if deps, ok := g.edges[node]; ok {
		return deps
	}
	return nil
}

// Validate traverses the validation graph
func (g *Graph) Validate(num int) []string {
	errors := make([]string, 0)
	if len(g.nodes) == 0 {
		return errors
	}

	stack := stack.New()
	stack.Push(g.nodes[0])
	seen := make(map[*Node]bool)
	depFailed := make(map[*Node]bool)

	for stack.Len() != 0 {
		val := stack.Pop()
		if node, ok := val.(*Node); ok {
			if _, ok := seen[node]; !ok {
				failed := false
				if depFail, ok := depFailed[node]; !ok || !depFail {
					if !node.Validate(num) {
						errors = append(errors, node.Msg)
						failed = true
					}
				}
				seen[node] = true
				for _, dep := range g.edges[node] {
					stack.Push(dep)
					depFailed[dep] = failed
				}
			}
		}
	}

	return errors
}
