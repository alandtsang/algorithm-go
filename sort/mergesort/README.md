# Merge Sort
![mergesort](https://raw.githubusercontent.com/alandtsang/algorithm-go/master/pic/mergesort.gif)

```
Input:  [6 5 3 1 8 7 2 4]
Divide: [6 5 3 1] [8 7 2 4]
Divide: [6 5] [3 1] [8 7] [2 4]
Divide: [6] [5] [3] [1] [8] [7] [2] [4]

Merge:  [5 6] [1 3] [7 8] [2 4]
Merge:  [1 3 5 6] [2 4 7 8]
Merge:  [1 2 3 4 5 6 7 8]
```

English:

Merge Sort is an algorithm belonging to the “divide and conquer” class of algorithms.
- If the array has only one element, then return.
- Split the array into two equal subarrays.
- Sort separate subarrays.
- Merge sorted arrays.

Complexity
- Runtime: O(nlogn)
- Memory: O(n)

中文：

归并排序是一种属于分治法的算法。
- 如果数组只有一个元素，则返回
- 将数组尽量分为两个相等的子数组
- 对分开的子数组进行排序
- 将排序后的数组合并

复杂度
- 时间：O(nlogn)
- 空间：O(n)

