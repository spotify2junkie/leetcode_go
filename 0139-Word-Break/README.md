# Notes 
- 原先思想：iterate through worddict, 然后从开始match每一个字符如果第一个字符有匹配，字符串就减少但这样会有O(n^2*m)的复杂度
- leetcode思想: dp[i] 表示i 点的状态， 如果无重复 dp[i-len(key)] 应该是true  （ 记得加个初始状态) 