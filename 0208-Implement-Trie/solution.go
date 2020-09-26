type Trie struct {
	value byte
	c     [26]*Trie
	end   int
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	s := len(word)
	for i := 0; i < s; i++ {
		idx := int(word[i] - 'a')
		if this.c[idx] == nil {
			this.c[idx] = &Trie{value: word[i]}
		}
		this = this.c[idx]
	}
	this.end++
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	size := len(word)
	for i := 0; i < size; i++ {
		idx := int(word[i] - 'a')
		if this.c[idx] == nil {
			return false
		}
		this = this.c[idx]
	}
	if this.end > 0 {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	size := len(prefix)
	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'  //define 开头
		if this.c[idx] == nil { // if any happen
			return false
		}
		this = this.c[idx]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */