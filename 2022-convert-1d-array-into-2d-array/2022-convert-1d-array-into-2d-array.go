func construct2DArray(original []int, m int, n int) [][]int {
    if len(original) != m*n {
        return [][]int{}
    }

    result := make([][]int, 0, m)
    for i := 0; i < m; i++ {
        result = append(result, original[i*n:(i+1)*n])
    }

    return result
}