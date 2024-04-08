func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1
    var res []int
    fmt.Println(numbers[left], numbers[right])
    for left < right{
        if (numbers[left] + numbers[right] > target) {
            right -= 1
        }else if (numbers[left] + numbers[right] < target){
            left += 1
        } else{
            res = append(res, left+1)
            res = append(res, right+1)
            return res
        }
    }
    return []int{}    
}