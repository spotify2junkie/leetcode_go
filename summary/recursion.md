## Recursion 总结
递归最重要的两个特征： 结束条件和自我调用。 

```
int func(你今年几岁) {
    // 最简子问题，结束条件
    if (你1999年几岁) return 我0岁;
    // 自我调用，缩小规模
    return func(你去年几岁) + 1;   
}

```

**一点心得**：明白一个函数的作用并相信它能完成这个任务，千万不要试图跳进细节。

```
/* 看不懂没关系，底下有更详细的分析版本，这里突出体现递归的简洁优美 */
int pathSum(TreeNode root, int sum) {
    if (root == null) return 0;
    return count(root, sum) + 
        pathSum(root.left, sum) + pathSum(root.right, sum);
}
int count(TreeNode node, int sum) {
    if (node == null) return 0;
    return (node.val == sum) + 
        count(node.left, sum - node.val) + count(node.right, sum - node.val);
}
```

**分治**

分治算法可以分三步走：分解 -> 解决 -> 合并

如mergesort 

```
void merge_sort(一个数组) {
    if (可以很容易处理) return;
    merge_sort(左半个数组);
    merge_sort(右半个数组);
    merge(左半个数组, 右半个数组);
}
```

### leetcode 0784 总结 



### leetcode 0050 总结

- 摆明结束条件
- 然后再自我调用
- 分情况解决



### leetcode 0104 总结

- 找到`root == 0`的情,depth 是 1 
- 相信一个func 可以解决



### leetcode 464 总结