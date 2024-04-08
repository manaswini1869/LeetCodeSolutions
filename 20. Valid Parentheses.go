func isValid(s string) bool {
    brac_map := map[rune]rune{'}': '{', ']': '[', ')': '('}
    var brac_track []rune

    for _,char := range s {
        if _, ok := brac_map[char]; !ok {
            brac_track = append(brac_track, char)
        } else {
            if len(brac_track) == 0 || brac_track[len(brac_track)-1] != brac_map[char] {
                return false
            } 
            brac_track = brac_track[:len(brac_track)-1]
        }
    }
    fmt.Println(brac_track)

    return len(brac_track) == 0
    
}