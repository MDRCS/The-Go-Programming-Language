package main

import (
	"fmt"
)

//var graph = make(map[string]map[string]bool) // key : string ; value : map(string,bool)
//// example -> [[A,[B,true],[C,false]],[B,[A,false],[C,true]]] ..

var graph = map[string]map[string]bool{
	"A" : {"B" : false, "C" : true},
	"B" : {"A" : false, "C" : true},
	"C" : {"B" : false, "A" : false},
}


func main() {


	fmt.Println(graph)
	add_edge("A","B")
	fmt.Println(graph)

}

func add_edge(from, to string) {
	//graph[from][to] = true

	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}

	edges[to] = true
}

func hasEdge(from,to string) bool{
	return graph[from][to]
}
