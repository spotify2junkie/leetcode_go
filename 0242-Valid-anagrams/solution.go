func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false 
    }
    sm := make(map[byte]int)
    for i := range s {
        sm[s[i]]++
        sm[t[i]]--
    }

    for _, n := range sm {
		if n != 0 {
			return false
		}
	}

	return true
}
