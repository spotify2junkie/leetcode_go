func findWords(board [][]byte, words []string) []string {
        var results []string

        m := len(board)
        if m == 0 {
                return results
        }

        n := len(board[0])
        if n == 0 {
                return results
        }

        trie := buildTrie(words) //nnn build trie tree 

        for i := 0; i < m; i++ {
                for j := 0; j < n; j++ {
                        dfs(board, i, j, trie, &results)
                }
        }

        return results
}

func dfs(board [][]byte, i, j int, trie *TrieNode, result *[]string) {
    c := board[i][j]
    if c == '#' || trie.next[int(c - 'a')] == nil {
        return 
    }

    trie = trie.next[int(c-'a')]
    if len(trie.word) > 0 {
        *result = append(*result,trie.word)
        trie.word = ""
    }
    board[i][j] = '#'
    if i > 0 {
		dfs(board, i-1, j, trie, result)
	}

	if i < len(board)-1 {
		dfs(board, i+1, j, trie, result)
	}

	if j > 0 {
		dfs(board, i, j-1, trie, result)
	}

	if j < len(board[0])-1 {
		dfs(board, i, j+1, trie, result)
	}

	board[i][j] = c
}



func buildTrie(words []string) *TrieNode {
	root := new(TrieNode)
	for _, word := range words {
		cur := root
		for _, c := range word {
			cidx := int(c - 'a')
			if cur.next[cidx] == nil {
				cur.next[cidx] = new(TrieNode) //nnn give an new node
			}
			cur = cur.next[cidx]
		}
		cur.word = word
	}
	return root //nnn where is root
}

type TrieNode struct {
	next [26]*TrieNode
	word string
}