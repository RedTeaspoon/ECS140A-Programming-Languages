package lgraph

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
func arrContains(arr []rune, label rune)bool{
	for _, edgeLabel := range arr{
		if edgeLabel == label{
			return true
		}
	}
	return false
}

func findTarget(g1, g2 LGraph, target node, lenLeft uint, currentEdge edge, visited []rune) ([]rune, bool){
	// add edge to sequence
	visited = append(visited, currentEdge.label)

	// check if current edge leads to target node
	if lenLeft >= 0 && currentEdge.destination == target{
		return visited, true
	// check if adding edge exceeds length limit
	} else if lenLeft < 0{
		return nil, false
	}

	// get all outgoing edges from the given node
	edges, _ := g1(currentEdge.destination)
	for _, edge := range edges{
		// check that edge isn't in visited
		if ! arrContains(visited, edge.label){
			visited, pathExists := findTarget(g1, g2, target, lenLeft-1, edge, visited)
			if pathExists{
				return visited, pathExists
			}
		}
	}
	
}

// FindSequence returns (S, true) if there is a sequence S of length k from node
// s to node t in graph g1 and S is not a sequence from s to t in graph g2; else
// it returns (nil, false).
func FindSequence(g1, g2 LGraph, s, t node, k uint) ([]rune, bool) {
	// TODO: Complete the function.
	panic("TODO: implement this!")
	var visited = make([]rune, 0) 
	var seqExists bool
	g1StartEdges, g1StartExists := g1(s)
	_, g1EndExists := g1(t)
	
	g2StartEdges, g2StartExists := g2(t)
	_, g2EndExists := g2(t)

	// check if start & end node are the same + path is at least 0
	if(k >= 0 && s == t && g1StartExists && ! g2StartExists){
		return visited, true;
	// check that both start and end nodes exist & path length is long enough
	} else if k >= 1 && g1StartExists && g1EndExists{
		// for every edge from start
		for _, edge := range g1StartEdges{
			visited, seqExists = findTarget(g1, g2, t, k-1, edge, visited)
			if seqExists{
				break
			}
		}
		if seqExists{

		}
	// need to check that case exists for this
	}
	return nil, false;
}
