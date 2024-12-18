func singleNumber(nums []int) int {
    singleNumber := 0

    for _, val := range nums {
        singleNumber ^= val
    }

    return singleNumber
}