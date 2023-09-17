package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// This app is a time scheduler for a given set of tasks. It
// is meant to be a restful api that can be used by a front end
// to schedule tasks. It takes in a set of tasks, and prints out the
// shortest time to complete all tasks.

// The graph is a directed acyclic graph, the edges are weighted by the
// time between tasks, and the nodes are the tasks themselves.

type Graph struct {
	Nodes []Node `json:"Nodes"`
	Edges [][2]Node
}

type Node struct {
	Name     string `json:"Name"`
	Id       int    `json:"Id"`
	Day      int    `json:"Day"`
	Time     int    `json:"Time"`
	Priority int    `json:"Priority"`
}

func (g *Graph) AddNode(n Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1 Node, n2 Node) {
	g.Edges = append(g.Edges, [2]Node{n1, n2})
}

func main() {

	// Populate the graph with the nodes in data.json
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	var graph Graph

	err = json.Unmarshal(data, &graph.Nodes)

	if err != nil {
		fmt.Println(err)
	}

	// Print out the graph
	fmt.Println(graph.Nodes[1].Day)

}
