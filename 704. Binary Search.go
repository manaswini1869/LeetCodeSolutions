func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    var mid int
    for left <= right {
        mid = int((left+right) / 2)
        fmt.Println(mid)
        if nums[mid] == target {
            return mid
        } else if nums[left] < target {
            left += 1
        }else{
            right -= 1
        }
        
    }  
    return -1  
}