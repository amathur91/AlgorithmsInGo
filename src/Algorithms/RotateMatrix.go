package algorithms
/**
Problem:https://leetcode.com/problems/rotate-image
Results :
21 / 21 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 2.3 MB

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate Image.
Memory Usage: 2.3 MB, less than 49.47% of Go online submissions for Rotate Image.
 */
//func main() {
//	matrix := [][]int {{1,2,3},
//						{4,5,6},
//						{7,8,9}}
//
//	rotateMatrix(matrix)
//}

type MatrixPoints struct {
	rowIndex int
	columnIndex int
}

func rotateMatrix(matrix [][]int) {
		if len(matrix) > 0 &&  len(matrix[0]) > 0 {
			topLeftPoint := MatrixPoints{rowIndex: 0, columnIndex: 0}
			topRightPoint := MatrixPoints{rowIndex: 0, columnIndex: len(matrix[0]) - 1}
			leftCornerPoint := MatrixPoints{rowIndex: len(matrix) - 1, columnIndex: 0}
			rightCornerPoint := MatrixPoints{rowIndex: len(matrix) - 1, columnIndex: len(matrix[0]) - 1}

			for ; topLeftPoint.columnIndex < topRightPoint.columnIndex && leftCornerPoint.columnIndex < rightCornerPoint.columnIndex; {
				rightRowData := copyTopRowToRight(&matrix, &topLeftPoint, &topRightPoint, &topRightPoint, &rightCornerPoint)
				bottomRowData := copyRightRowToBottom(&rightRowData, &rightCornerPoint, &leftCornerPoint, &matrix)
				leftRowData := copyBottomRowToLeft(&bottomRowData, &matrix, &leftCornerPoint, &topLeftPoint)
				copyLeftRowToTop(&leftRowData, &topLeftPoint, &topRightPoint, &matrix)
				topLeftPoint.rowIndex++
				topLeftPoint.columnIndex++
				topRightPoint.rowIndex++
				topRightPoint.columnIndex--
				leftCornerPoint.rowIndex--
				leftCornerPoint.columnIndex++
				rightCornerPoint.rowIndex--
				rightCornerPoint.columnIndex--
			}
		}
}

func copyLeftRowToTop(leftRowData *[]int,startPoint *MatrixPoints,endPoint *MatrixPoints,matrix *[][]int) {
	dataPointIndex := 0
	for columnIndex := startPoint.columnIndex; columnIndex <= endPoint.columnIndex; columnIndex++ {
		(*matrix)[startPoint.rowIndex][columnIndex] = (*leftRowData)[dataPointIndex]
		dataPointIndex++
	}
}

func copyBottomRowToLeft(bottomRowData *[]int, matrix *[][]int, startPoint *MatrixPoints, endPoint *MatrixPoints) []int {
	returnData := make([]int, startPoint.rowIndex -  endPoint.rowIndex  + 1)
	dataIndex := 0
	returnData[dataIndex] = (*bottomRowData)[len(*bottomRowData) - 1]
	dataIndex++
	for rowIndex := startPoint.rowIndex - 1; rowIndex >= endPoint.rowIndex; rowIndex-- {
		returnData[dataIndex] = (*matrix)[rowIndex][startPoint.columnIndex]
		(*matrix)[rowIndex][startPoint.columnIndex] =(*bottomRowData)[dataIndex]
		dataIndex++
	}
	return returnData
}

func copyRightRowToBottom(rightRowData *[]int, startPoint *MatrixPoints, endPoint *MatrixPoints, matrix *[][]int) []int {
	returnData := make([]int, (startPoint.columnIndex - endPoint.columnIndex) + 1)
	//copy bottom row to returnData
	dataIndex := 0
	returnData[dataIndex] = (*rightRowData)[len(*rightRowData) - 1]
	dataIndex++
	for columnIndex := startPoint.columnIndex - 1; columnIndex >= endPoint.columnIndex; columnIndex-- {
		returnData[dataIndex] = (*matrix)[startPoint.rowIndex][columnIndex]
		(*matrix)[startPoint.rowIndex][columnIndex] = (*rightRowData)[dataIndex]
		dataIndex++
	}
	return returnData
}

func copyTopRowToRight(matrix *[][]int, startLeftPoint *MatrixPoints, endRightPoint *MatrixPoints, topRightPoint *MatrixPoints, rightCornerPoint *MatrixPoints) []int{
	//back up of right side
	returnData := make([]int, rightCornerPoint.rowIndex - topRightPoint.rowIndex + 1)
	tempData := make([]int, endRightPoint.columnIndex - startLeftPoint.columnIndex + 1)
	tempDataIndex := 0
	for columnIndex:= startLeftPoint.columnIndex; columnIndex <= topRightPoint.columnIndex; columnIndex++ {
		tempData[tempDataIndex] = (*matrix)[startLeftPoint.rowIndex][columnIndex]
		tempDataIndex++
	}

	tempDataIndex = 0
	for rowIndex := topRightPoint.rowIndex; rowIndex <= rightCornerPoint.rowIndex; rowIndex++ {
		returnData[tempDataIndex] = (*matrix)[rowIndex][topRightPoint.columnIndex]
		(*matrix)[rowIndex][topRightPoint.columnIndex] = tempData[tempDataIndex]
		tempDataIndex++
	}
	return returnData
}