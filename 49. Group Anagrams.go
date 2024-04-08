func groupAnagrams(strs []string) [][]string {
    res_map := make(map[string][]string)
    for i:=range(len(strs)){
        str_bytes := []byte(strs[i])
        sort.Slice(str_bytes, func(i int, j int) bool {
            return str_bytes[i] < str_bytes[j]
        })
        fmt.Println(string(str_bytes))
        sorted_str := string(str_bytes)
        res_map[sorted_str] = append(res_map[sorted_str], strs[i])
    }
    fmt.Println(res_map)
    result := make([][]string, 0)
    for _, value := range res_map {
        result = append(result, value)
    }
    return result
    
}