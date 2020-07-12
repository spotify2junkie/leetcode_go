## Notes from Leetcode

- `Target` - the value that you are searching for
- `Index` - the current location that you are searching
- `Left, Right` - the indicies from which we use to maintain our search Space
- `Mid` - the index that we use to apply a condition to determine if we should search left or right



> Binary Search can take many alternate forms and might not always be as straight forward as searching for a specific value. Sometimes you will have to apply a **specific condition** or rule to determine which side (left or right) to search next.



### Three steps of Binary Search 

1. ***Pre-processing\*** - Sort if collection is unsorted.
2. ***Binary Search\*** - Using a loop or recursion to divide search space in half after each comparison.
3. ***Post-processing\*** - Determine viable candidates in the remaining space.



### Three templates 

- template 1
  - Initial Condition:` left = 0, right = length-1`
  - Termination: `left > right`
  - Searching Left: `right = mid-1`
  - Searching Right: `left = mid+1`

- Template 3
  - 查找邻居
  - 有重复元素就使用这个



P.S : 

```python
## template 1
## 
def binarySearch(nums, target):
    """
    :type nums: List[int]
    :type target: int
    :rtype: int
    """
    if len(nums) == 0:
        return -1

    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1

    # End Condition: left > right
    return -1
  

```



### Notes on Leetcode 702 size 未知如何解题 

- size 未知怎么二分？ 



### Notes on start, end 

- Start 和 end 的起始位置要注意选用哪个， 
- 如果是数组的话可以选用 0, len(nums)-1 ; 如果是一个数字n的话还是选用 1, n .



### Notes on template 3 

- 大部分问题可以解决



### Summary 二分搜索核心四点要素

- 1、初始化：start=0、end=len-1
- 2、循环退出条件：start + 1 < end
- 3、比较中点和目标值：A[mid] ==、 <、> target
- 4、判断最后两个元素是否符合：A[start]、A[end] ? target



## Template 3 和 1 的比较 

- template 3 `start + 1 < end` 
- Templaote 1 `start <= end`
- template 1 :` left = mid + 1 , right = mid -1` ; template 3 `left = mid , right = mid`

