func containsDuplicate(nums []int) bool {
    res := make(map[int]int)
    for i:= range(len(nums)) {
        if _,ok := res[nums[i]]; ok{
            res[nums[i]] += 1
        } else {
            res[nums[i]] += 1
        }
        if res[nums[i]] >= 2 {
            return true
        }
    }    
    return false
}