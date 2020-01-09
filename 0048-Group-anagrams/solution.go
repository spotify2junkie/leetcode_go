// hash table , play with byte 

func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string) // map[string][]string
    for _,str := range strs {
        b := []byte(str)
        sort.Slice(b, func(i, j int) bool { // sort byte 
            return b[i] < b[j]
        })
        key := string(b)
        groups[key] = append(groups[key],str)
    }

    ret := make([][]string,0,len(groups)) // 容量和长度的问题
    for _,v := range groups {
        ret = append(ret,v)
    }
    return ret 
}