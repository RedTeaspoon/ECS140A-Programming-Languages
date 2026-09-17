package lgraph
import "fmt"

type node uint

type edge struct {
	destination node
	label       rune
}

// LGraph is a function representing a directed labeled graph. If the node exists
// in the graph, the function returns true along with the set of outgoing edges
// from that node, otherwise false and nil.
type LGraph func(node) ([]edge, bool)

// check if sequence contains rune label
func labelsContainsEdge(labels []rune, edge edge)bool{
	for _, label := range labels{
		if label == edge.label{
			return true
		}
	}
	return false
}

func findTarget(g1 LGraph, target node, lenLeft int, currentEdge edge, visited []rune, visitedArr *[][]rune){
	// add edge to sequence
	visited = append(visited, currentEdge.label)

	// check if current edge leads to target node
	if lenLeft == 0 && currentEdge.destination == target{
		temp := make([]rune, len(visited))
		copy(temp, visited[:])
		*visitedArr = append(*visitedArr, temp)
		return
	// check if adding edge exceeds length limit
	} else if currentEdge.destination == target || lenLeft < 0{
		return
	}
	
	// get all outgoing edges from the given node (from currentEdge)
	edges, _ := g1(currentEdge.destination)
	for _, edge := range edges{
		// fmt.Println("checking edges")
		// check that edge isn't in visited
		if ! labelsContainsEdge(visited, edge){
			// fmt.Println("recursive call")
			findTarget(g1, target, lenLeft-1, edge, visited, visitedArr)
		}
	}
}

// check if sequence contains rune label, returns index of edge containing label
func edgesContainsLabel(edges []edge, label rune)int{
	for index, edge := range edges{
		if edge.label == label{
			return index
		}
	}
	return -1
}

// FindSequence returns (S, true) if there is a sequence S of length k from node
// s to node t in graph g1 and S is not a sequence from s to t in graph g2; else
// it returns (nil, false).
func FindSequence(g1, g2 LGraph, s, t node, k uint) ([]rune, bool) {
	var visited = make([]rune, 0)
	var visitedArr = make([][]rune, 0)

	g1StartEdges, g1StartExists := g1(s)
	_, g1EndExists := g1(t)
	
	_, g2StartExists := g2(s)
	_, g2EndExists := g2(t)

	// check if start & end node are the same + path is at least 0
	// or nodes in g1 don't exist in g2
	if s == t && g1StartExists && ! g2StartExists {
		return visited, true;
	// check that both start and end nodes exist & path length is long enough
	} else if k >= 1 && g1StartExists && g1EndExists{
		// for every edge from start
		for _, edge := range g1StartEdges{
			findTarget(g1, t, int(k-1), edge, visited, &visitedArr)
		}

		// check if g1 has any valid sequences
		if len(visitedArr) == 0{
			fmt.Println("No valid sequences")
			return nil, false
		}

		fmt.Println("Checking visited")
		for _, visited := range visitedArr{
			for _, label := range visited{
				fmt.Print(string(label))
			}
			fmt.Println()
		}

		// sequence exists & either start or end of g2 doesn't exist
		if !(g2StartExists && g2EndExists){
			fmt.Println("return 5")
			return visitedArr[0], true;
		}else{
			for _, checkingPath := range visitedArr{
				g2Edges, _ := g2(s)
				for index, label := range checkingPath{
					// fmt.Println(string(label))
					// fmt.Println(string(g2Edges[0].label))
					nextNodeIndex := edgesContainsLabel(g2Edges, label)
					if nextNodeIndex == -1{
						fmt.Println("Failed to find next node")
						return checkingPath, true
					}else if index == len(checkingPath) - 1 && nextNodeIndex >= 0 && g2Edges[nextNodeIndex].destination != t{
						fmt.Println("new if")
						return checkingPath, true
					}else{
						g2Edges, _ = g2(g2Edges[nextNodeIndex].destination)
					}
					
					
				}
			}
		}
	}
	fmt.Println("last return")
	return nil, false;
}
