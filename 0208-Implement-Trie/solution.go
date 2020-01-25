package mains

type Trie struct {
	value   byte //define byte for slice purpose
	childen [26]*Trie
	end     int
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a' // unicode expression can get rune , value: a-z
		if node.childen[idx] == nil {
			node.childen[idx] = &Trie{value: word[i]} //get address for value
		}
		node = node.childen[idx]

	}

	node.end++
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a' // unicode expression
		if node.childen[idx] == nil {
			return false
		}
		node = node.childen[idx]

	}

	if node.end > 0 {
		return true
	}

	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	size := len(prefix)
	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'        //define 开头
		if node.childen[idx] == nil { // if any happen
			return false
		}
		node = node.childen[idx]
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
