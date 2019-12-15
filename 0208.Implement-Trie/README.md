# Prefix Tree
利用空间替代时间

’‘’
type Trie struct {
    value byte //define byte for slice purpose
    childen [26]*Trie
    end int 
}
‘’‘

# 思路
- 给每个字母一个位置（node)，一个单词的不同字母通过指针相连
- `insert` 方法和`search` 方法可以类似实现，所不同的地方在于insert获取&Trie地址以及给val赋值（node赋值）, search 是bool值的check
- `startwith` 和`search`是一样的

# 改进
- 多个node-pass之后就又传给了node，有点消耗内存（这种存变量行为），是否可指针实现？