# Notes 
- First thought: sliding window but cant 

# 其他人的思路
状态转移方程是现在的最大或者最小乘multiply `nums[i]`

- how to get max `max=max(nums[i],max*nums[i])`
- how to get minimum `min=min(nums[i],min*nums[i])`

其实刚开始我也是想着dp，但是没想到怎么构建dp。我的思路是去找有多少个负值，但是可以通过dp得到当前的最大最小值，然后乘以current的值，如果当前是负值那当前就乘以之前的最小值，争执就乘以之前累积下来的最大值。