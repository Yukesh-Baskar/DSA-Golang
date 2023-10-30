package graph

import "fmt"

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex, toVertex := g.getVertex(from), g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Printf("vertex not found. fromVertex: %v toVertex: %v \n", fromVertex, toVertex)
		return
	} else if g.contains(fromVertex, toVertex) {
		fmt.Printf(" \n edge already exist from : %v and to : %v \n", fromVertex.Key, toVertex.Key)
		return
	}

	fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
}

func (g *Graph) contains(from, to *Vertex) bool {
	if from.Adjacent == nil {
		return false
	} else {
		for _, v := range from.Adjacent {
			if v.Key == to.Key {
				return true
			}
		}
	}
	return false
}

func (g *Graph) AddVertex(key int) {
	if isVertexExist(g.Vertices, key) {
		fmt.Println(" vertex not added, key already exist", key)
		return
	}
	g.Vertices = append(g.Vertices, &Vertex{Key: key})
}

func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.Vertices {
		if v.Key == key {
			return g.Vertices[i]
		}
	}
	return nil
}

func (g *Graph) PrintVertex() {
	for _, v := range g.Vertices {
		fmt.Printf("\n vertex: %v ", v.Key)
		for _, v := range v.Adjacent {
			fmt.Print(" edge: ", v.Key)
		}
	}
}

func isVertexExist(vv []*Vertex, key int) bool {
	for _, v := range vv {
		if v.Key == key {
			return true
		}
	}
	return false
}
