func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])

    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    maxSide := 0

    for i := 0; i < m; i++ {
        onesCount := 0
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                onesCount++
            }
        }

        if onesCount == 0 {
            continue
        }

        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                if i == 0 || j == 0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
                }

                if dp[i][j] > maxSide {
                    maxSide = dp[i][j]
                }
            }
        }
    }

    return maxSide * maxSide
}

func min(a, b, c int) int {
    if a < b {
        if a < c {
            return a
        }
        return c
    }
    if b < c {
        return b
    }
    return c
}
