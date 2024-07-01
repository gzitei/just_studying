# 1550. Three Consecutive Odds
##### 📌 Algorithms | 📆 2024-07-01 | 🟢 Easy | 📊 67.94% | 🌐 [Leet-Code #1550](https://leetcode.com/problems/three-consecutive-odds)
---
```yaml
🔖tags: #array
```
---
Given an integer array `arr`, return `true` if there are three consecutive odd numbers in the array. Otherwise, return `false`.

**Example 1:**

```
Input: arr = [2,6,4,1]
Output: false
Explanation: There are no three consecutive odds.

```

**Example 2:**

```
Input: arr = [1,2,34,3,4,5,7,23,12]
Output: true
Explanation: [5,7,23] are three consecutive odds.

```

**Constraints:**

- `1 <= arr.length <= 1000`
- `1 <= arr[i] <= 1000`


---
> ### Dicas:
>💡Check every three consecutive numbers in the array for parity.