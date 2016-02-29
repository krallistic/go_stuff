package main

import (
	"fmt"
)

/*
1
4 6
1 2 2
1 3 4
1 4 8
2 3 3
2 4 3
3 4 1
*/

type graph struct {
	nodes []node
	Edges []edge
}

func (g graph) addEdge(a, b, d int) {
	var e edge
	var node_a, node_b node
	node_a.id = a
	node_b.id = b
	g.addNode(node_a)
	g.addNode(node_b)
	e.length = d
	e.end_1 = g.nodes[a]
	e.end_2 = g.nodes[b]
	g.Edges = append(g.Edges, e)
	g.nodes[a].edges = append(g.nodes[a].edges, e)
	g.nodes[b].edges = append(g.nodes[b].edges, e)
}

func (g graph) addNode(n node) {
	if g.nodes[n.id].id == 0 {
		g.nodes[n.id] = n
	}
}

func (g graph) getNote(index int) node {
	return g.nodes[index]
}

type node struct {
	id    int
	edges []edge
}

func (n node) getNeigbors() []node {
	var neigbors []node
	neigbors = make([]node, len(n.edges))
	for _, e := range n.edges {
		if e.end_1.id == n.id {
			neigbors = append(neigbors, e.end_2)
		} else {
			neigbors = append(neigbors, e.end_1)
		}
	}

	return neigbors
}

type edge struct {
	length int
	end_1  node
	end_2  node
}

func (e edge) get_otherside(n node) node {
	if e.end_1.id == n.id {
		return e.end_2
	}
	if e.end_2.id == n.id {
		return e.end_1
	}
	panic("No should be on one Side")
}

func get_n_shortest_route(g graph) int {
	var current_min int
	current_min = 999999999

	for _, n := range g.nodes {
		//fmt.Println(n.id)
		for _, e := range n.edges {
			//get Node
			second_Node := e.get_otherside(n)
			if e.length <= current_min {
				for _, e2 := range second_Node.edges {
					third_node := e2.get_otherside(second_Node)
					if e.length+e2.length <= current_min && n.id != third_node.id {
						current_min = e.length + e2.length
					}
				}
			}

		}
	}

	return current_min
}

type middleNode struct {
	min       int
	secondMin int
}

func (m *middleNode) update(d int) {
	fmt.Println("Update", m, d)
	if (m.min == 0) && (m.secondMin == 0) {
		m.min = d
		return
	}
	if m.secondMin == 0 {
		m.secondMin = d
		return
	}
	if (m.min == 0) || (d < m.min) {
		m.min = d
		return
	}
	if (m.secondMin == 0) || (d > m.min && d < m.secondMin) {
		m.secondMin = d
	}
}

func (m *middleNode) sum() int {
	if m.min != 0 && m.secondMin != 0 {
		return m.min + m.secondMin
	}
	return 99999

}

func getMinMiddle(nodes []middleNode) int {
	minimum := 999999
	for _, nm := range nodes {
		if minimum >= nm.sum() {
			minimum = nm.sum()
		}
	}
	return minimum
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)
		nodes := make([]middleNode, n+1)
		for j := 0; j < m; j++ {
			var a, b, d int
			fmt.Scan(&a, &b, &d)
			nodes[a].update(d)
			nodes[b].update(d)

		}
		fmt.Println(nodes)
		fmt.Println(getMinMiddle(nodes))

	}
}
