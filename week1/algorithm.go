package week1

/**
 * leetcode 48. 旋转图像
 * 给定一个 n×n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
 * 你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
 */

// Rotate1 用了另一矩阵，不符题意
func Rotate1(matrix [][]int) {
	rotate := make([][]int, len(matrix))
	rank := len(matrix)
	for i := 0; i < rank; i++ {
		rotate[i] = make([]int, rank)
	}
	for i := 0; i < rank; i++ {
		for j := 0; j < rank; j++ {
			rotate[i][j] = matrix[rank-1-j][i]
		}
	}
	for i := 0; i < rank; i++ {
		copy(matrix[i], rotate[i])
	}
}

// Rotate2 找规律，矩阵是正方形，有4条边，所以每次转动4个
func Rotate2(matrix [][]int) {
	rank := len(matrix)
	for i := 0; i < rank-1-i; i++ {
		for j := i; j < rank-1-i; j++ {
			//leftUp := matrix[i][j]
			//rightUp := matrix[j][rank-1-i]
			//rightDown := matrix[rank-1-i][rank-1-j]
			//leftDown := matrix[rank-1-j][i]

			//matrix[j][rank-1-i] = leftUp
			//matrix[rank-1-i][rank-1-j] = rightUp
			//matrix[rank-1-j][i] = rightDown
			//matrix[i][j] = leftDown

			matrix[j][rank-1-i], matrix[rank-1-i][rank-1-j], matrix[rank-1-j][i], matrix[i][j] = matrix[i][j], matrix[j][rank-1-i], matrix[rank-1-i][rank-1-j], matrix[rank-1-j][i]
		}
	}
}
