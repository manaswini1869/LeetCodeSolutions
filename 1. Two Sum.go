func twoSum(nums []int, target int) []int {
    res_hash := make(map[int]int)
    for i := range(len(nums)){
        req := target - nums[i]
        if _, ok := res_hash[req]; ok{
            return []int{res_hash[req], i}
        }else {
            res_hash[nums[i]] = i
        }
    }
    return []int{}
}