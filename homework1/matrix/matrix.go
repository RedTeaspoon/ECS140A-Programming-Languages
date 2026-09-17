package matrix

// If needed, you may define helper functions here.

// AreAdjacent returns true if a and b are adjacent in lst.
func AreAdjacent(a, b int, lst []int) bool {
	for i, value := range lst{
		if value == a && ((i-1 >= 0 && lst[i-1] == b) || (i+1 < len(lst) && lst[i+1] == b)){
			return true
		}
	}
	return false
}

// Transpose returns the transpose of the 2D matrix mat.
func Transpose(mat [][]int) [][]int {
	if len(mat) == 0{
		return mat
	}

	// assume matrix is square
	transMat := make([][]int, len(mat[0]))
	for i, _ := range transMat{
		transMat[i] = make([]int, len(mat))
	}

	for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[i]); j++ {
            transMat[j][i] = mat[i][j]
        }
    }
	return transMat
}

// AreNeighbors returns true if a and b are neighbors in the 2D matrix mat.
func AreNeighbors(mat [][]int, a, b int) bool {
	for i, row := range mat{
		for j, value := range row{
			if value == a && ((i-1 >= 0 && mat[i-1][j] == b) || (i+1 < len(mat) && mat[i+1][j] == b) ||
				(j-1 >= 0 && mat[i][j-1] == b) || (j+1 < len(row) && mat[i][j+1] == b)){
				return true
			}
		}
	}
	return false
}
