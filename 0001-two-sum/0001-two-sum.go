func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        result := target - nums[i]
        if _, ok := indexMap[result]; ok {
            return []int{i,indexMap[result]}
        }
        indexMap[nums[i]] = i
    }
    return nil
}