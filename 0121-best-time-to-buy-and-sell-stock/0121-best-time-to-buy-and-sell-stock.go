func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    buy := prices[0]
    profit := 0

    for _, val := range prices {
        if val < buy {
            buy = val
        } else if val - buy > profit {
            profit = val - buy
        }
    }

    return profit
}