func evalRPN(tokens []string) int {
    res := []int{}
    var result int
    operations := map[string]bool {
        "*": true,
        "+": true,
        "-": true,
        "/": true,
    }
    for _, token := range(tokens) {
        if operations[token] && len(res) > 0 {
            a := res[len(res)-1]
            b := res[len(res)-2]
            if token == "+"{
                result = a+b
            }else if  token == "-" {
                result = b-a
            }else if  token == "*" {
                result = b*a
            }else if  token == "/" {
                result = int(b/a)
            }
            res = res[:len(res)-2]
            res = append(res, result)
        } else {
            num, _ := strconv.Atoi(token)
            res = append(res, num)
        }
    }
    return res[0]
    
}