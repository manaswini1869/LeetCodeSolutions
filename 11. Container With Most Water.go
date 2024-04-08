func maxArea(height []int) int {
    left, right := 0, len(height)-1
    res := 0
    var h int
    for left < right {
        if height[left] < height[right] {
            h = height[left]
        }else{
            h = height[right]
        }
        w := right - left
        curr_max := h * w
        if curr_max > res {
            res = curr_max
        }
        if height[left] > height[right] {
            right -= 1
        }else {
            left += 1
        }
    }
    return res
    
}