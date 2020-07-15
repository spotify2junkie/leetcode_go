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

- [ ] to-do [link](https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/1216/)

  