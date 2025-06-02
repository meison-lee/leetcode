package main

func searchMatrix(matrix [][]int, target int) bool {    

    min, max := matrix[0][0], matrix[len(matrix) - 1][len(matrix[0]) - 1]
    if target < min || target > max{
        return false
    }

    if len(matrix) == 1{
        return search(matrix, target, 0)
    }

    start, end := 0, len(matrix)-1
    row := 0

    for start < end{        
        mid := start + (end - start)/2
        if matrix[mid][0] == target{
            return true
        }else if matrix[mid][0] > target{            
            end = mid
        }else {            
            if start == mid {                             
                break
            }
            start = mid
        }
    }

    if start == end {
        row = start
    }else if target < matrix[end][0]{
        row = start
    }else {
        row = end
    }    

    return search(matrix, target, row)
}

func search(matrix [][]int, target int, row int) bool {
    start, end := 0, len(matrix[row])-1
	for start <= end {
		mid := (start + end) / 2

		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

    return false
}
