# 136. Single Number
##### 📌 Algorithms | 📅 2024-08-08 | 🟢 Easy | 📊 73.72% | 🌐 [Leet-Code #136](https://leetcode.com/problems/single-number)
---
```yaml
🔖tags: #array #bit-manipulation
```
---
Given a **non-empty** array of integers `nums`, every element appears _twice_ except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

**Example 1:**

```
Input: nums = [2,2,1]
Output: 1

```

**Example 2:**

```
Input: nums = [4,1,2,1,2]
Output: 4

```

**Example 3:**

```
Input: nums = [1]
Output: 1

```

**Constraints:**

- `1 <= nums.length <= 3 * 104`
- `-3 * 104 <= nums[i] <= 3 * 104`
- Each element in the array appears twice except for one element which appears only once.


---
> ### Dicas:
>💡Think about the XOR (^) operator's property.