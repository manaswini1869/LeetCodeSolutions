func isPalindrome(s string) bool {
    format_func := func(r rune) rune {
        if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
            return -1
        }
        return unicode.ToLower(r)
    }
    s = strings.Map(format_func, s)

    i,j := 0, len(s)-1
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i += 1
        j -= 1
    }
    return true
    
}