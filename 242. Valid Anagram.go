func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    res_s := make(map[byte]int)
    res_t := make(map[byte]int)

    for i := range(len(s)){
        // if _, ok := res_s[s[i]]; ok {
        //     res_s[s[i]] += 1
        // }else{
        //     res_s[s[i]] = 1
        // }
        // if _, ok := res_t[t[i]]; ok {
        //     res_t[t[i]] += 1
        // }else{
        //     res_t[t[i]] = 1
        // }
        res_s[s[i]]++
        res_t[t[i]]++
    }
     for k, v := range res_s {
        if res_t[k] != v {
            return false
        }
    }

    return true
    
    
}