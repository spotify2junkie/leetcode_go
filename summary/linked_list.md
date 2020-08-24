## Notes On Leetcode Explore Sec.

There are **two types** of linked list: singly linked list and doubly linked list.

- [ ] Understand the structure of singly linked list and doubly linked list;
- [ ] Implement traversal, insertion, deletion in a singly or doubly linked list;
- [ ] Analyze the complexity of different operations in a singly or doubly linked list;
- [ ] Use two-pointer technique (fast-pointer-slow-pointer technique) in the linked list;
- [ ] Solve classic problems such as reverse a linked list;
- [ ] Analyze the complexity of the algorithms you designed;
- [ ] Accumulate experience in designing and debugging.



### Node

- value 
- a refereince field to the next **node** 



### Node 取值

- unlike array, 如果要取特定元素，现要traverse other element 
- time complexity: `O(N)`.
- that's because of `insert` and `delete` method 



### Insert 

- `O(1)`
- 对于中间node的插入，只要制定前面一个连接和cur 连接方式就可以。



## 解题思路

- 先check 有没有不对的地方 

  ```go
  if head == nil || head.Next == nil || head.Next.Next == nil {
          return false
      }
  ```




## 官方双指针模板

```go
slow, fast := &ListNode{}, &ListNode{}
for slow != nil && fast != nil && fast.next != nil {
  	slow = slow.Next
  	fast = fast.Next
  if slow == fast {
    	return true
  }
}
return false 
```

- 官方提供的tips:

  - **1. Always examine if the node is null before you call the next field.**
  - **2. Carefully define the end conditions of your loop.**
  - **3. In many cases, you need to track the previous node of the current node.**

  

