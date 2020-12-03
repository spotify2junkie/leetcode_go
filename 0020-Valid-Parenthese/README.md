# Problem 
应该是用栈解决括号匹配问题

# Notes
- 尝试用前一个字节和后一个字节match，但不行
- `rune`是一个type，但不能用map 取key，要用`string` 函数转。
- 实现前括号和后括号匹配，同时匹配完之后，**出栈**，再匹配下一个。（对称关系）

# Furthermore


