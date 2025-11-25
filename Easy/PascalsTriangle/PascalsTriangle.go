package PascalsTriangle

func Generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if i == 0 || j == len(triangle[i])-1 || j == 0 {
				triangle[i][j] = 1
				continue
			}
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}
