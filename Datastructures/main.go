package main

import (
	hashtable "data-structures-learnings/Hash-Table"
	"fmt"
)

type A interface {
	get()
}

type StructA struct{}
type StructB struct{}

type EmbeddedInterfaceStruct struct {
	A
}

func (s StructA) get() {
	fmt.Println("Hello A")
}

func (s StructB) get() {
	fmt.Println("Hello B")
}

func main() {
	// e := EmbeddedInterfaceStruct{StructA{}}
	// e.get()
	// e = EmbeddedInterfaceStruct{StructB{}}
	// e.get()
	// tree := bst.NewTree(100)
	// tree.Insert(tree.Root, 1)
	// tree.Insert(tree.Root, 101)
	// tree.Insert(tree.Root, 110)
	// fmt.Printf("%+v \n", tree.Root)
	// fmt.Println(tree.Search(tree.Root, 110))
	// g := &graph.Graph{}

	// for i := 1; i <= 5; i++ {
	// 	g.AddVertex(i)
	// }
	// g.AddEdge(1, 2)
	// g.AddEdge(2, 3)
	// g.AddEdge(2, 4)
	// g.AddEdge(2, 5)
	// g.AddEdge(2, 6)
	// g.PrintVertex()
	// g.AddEdge(1, 2)

	hashtable.NewHashTable()
	// fmt.Println(ht)
}
