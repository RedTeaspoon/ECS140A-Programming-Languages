package disjointset
import "fmt"

type Set struct{
	unionedSets		map[int]node
}

type node struct{
	rep			*int
	repping		[]int
}
// func (n node) String() string {
// 	var num int = *n.rep
//     return fmt.Sprintf("{Rep: %d}", num)
// }

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

// TODO: implement a type that satisfies the DisjointSet interface.
func (set Set) FindSet(num int)int{
	repNode, exists := set.unionedSets[num]
	if ! exists{
		return num
	}

	repValue := repNode.rep
	if repValue == nil{
		return num
	}else{
		return *repValue
	}
}

func (set Set) UnionSet(a, b int) int{
	// fmt.Println("Joining: ", a, ", ", b)
	aNode, aExists := set.unionedSets[a]
	bNode, bExists := set.unionedSets[b]

	// if a == b && !aExists{
	// 	return a
	// } else if a == b && aExists{
	// 	return *aNode.rep
	// }

	// a & b aren't part of a union, a becomes rep
	if !aExists && !bExists{
		var newANode, newBNode node 

		newANode.rep = nil
		newBNode.rep = &a
		arr := []int{b}

		newANode.repping = arr
		set.unionedSets[a] = newANode
		set.unionedSets[b] = newBNode
		return a
	// joining single value to set
	}else if (aExists && !bExists) || (!aExists && bExists){
		if aNode.rep != nil{
			// fmt.Println("Test use 2.1")
			a = *aNode.rep
			aNode, _ = set.unionedSets[a]
			// fmt.Println("New a: ", a)
		}
		if bNode.rep != nil{
			// fmt.Println("Test use 2.2")
			b = *bNode.rep
			bNode, _ = set.unionedSets[b]
			// fmt.Println("New b: ", b)
		}

		if aExists{
			// fmt.Println("Test use 2.3")
			var newBNode node
			aNode.repping = append(aNode.repping, b)
			newBNode.rep = &a
			set.unionedSets[b] = newBNode
			return a
		}else{
			// fmt.Println("Test use 2.4")
			var newANode node
			bNode.repping = append(bNode.repping, a)
			newANode.rep = &b
			set.unionedSets[a] = newANode
			return b
		}

	}else{
		if aNode.rep != nil{
			// fmt.Println("Test use 3.1")
			a = *aNode.rep
			aNode, _ = set.unionedSets[a]
			// fmt.Println("New a: ", a)
		}
		if bNode.rep != nil{ // not used
			fmt.Println("Test use 3.2")
			b = *bNode.rep
			bNode, _ = set.unionedSets[b]
			// fmt.Println("New b: ", b)
		}

		// merge b into a
		if len(aNode.repping) >= len(bNode.repping){
			// fmt.Println("Test use 3.3")
			arr := bNode.repping[:]
			copy(arr, bNode.repping[:])
			for _, value := range arr{
				temp := set.unionedSets[value]
				temp.rep = &a
				set.unionedSets[value] = temp
			}
			bNode.rep = &a
			bNode.repping = make([]int, 0)
			aNode.repping = append(aNode.repping, arr...)

			set.unionedSets[a] = aNode
			set.unionedSets[b] = bNode
			return a
		// merge a into b
		}else{ // not used
			fmt.Println("Test use 3.4")
			arr := aNode.repping[:]
			copy(arr, aNode.repping[:])
			for _, value := range arr{
				temp := set.unionedSets[value]
				temp.rep = &b
				set.unionedSets[value] = temp
			}
			aNode.rep = &b
			aNode.repping = make([]int, 0)
			bNode.repping = append(bNode.repping, arr...)

			set.unionedSets[a] = aNode
			set.unionedSets[b] = bNode
			return b
		}
	}
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	var set Set
	set.unionedSets = make(map[int]node)
	// set.UnionSet(1, 2)
	// set.UnionSet(1, 3)
	// set.UnionSet(4, 5)
	// fmt.Println("Checking: ", set.FindSet(4))
	// set.UnionSet(2, 5)
	// fmt.Println("Checking: ", set.FindSet(4))

	// set.UnionSet(1, 2)
	// set.UnionSet(3, 4)
	// set.UnionSet(4, 5)
	// set.UnionSet(2, 5)
	// fmt.Println("Set: ", set)

	// set.UnionSet(-1, 2)
	// set.UnionSet(-1, 3)
	// set.UnionSet(-1, 4)

	// set.FindSet(-2)
	// set.FindSet(4)
	// set.FindSet(-3)
	// set.FindSet(-1)
	// set.FindSet(-2)
	// set.FindSet(3)
	// set.FindSet(6)
	return set
}
