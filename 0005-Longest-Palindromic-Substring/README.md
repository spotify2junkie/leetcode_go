# Notes
- 从第i个位置开始查找，前后元素，直到长度抵达左边和右边的最小值
- 有两个方向，先从左边扫到endpoint，作为`e`
- 先找到中间相同的长度
- `b` 和`e`都是从i最先开始, e往右移，b往左移
- 动态更新，找到新的回文起始点会从新的b开始