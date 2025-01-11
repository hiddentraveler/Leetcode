// Source: https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/description/
// Author: Neo Orez
// Difficulty: Easy

/*
A string is **good** if there are no repeated characters.

Given a string `s`, return *the number of **good substrings** of length **three** in* `s`.

Note that if there are multiple occurrences of the same substring, every occurrence should be counted.

A **substring** is a contiguous sequence of characters in a string.

**Example 1:**

```
Input: s = "xyzzaz"
Output: 1
Explanation: There are 4 substrings of size 3: "xyz", "yzz", "zza", and "zaz".
The only good substring of length 3 is "xyz".
```

**Example 2:**

```
Input: s = "aababcabc"
Output: 4
Explanation: There are 7 substrings of size 3: "aab", "aba", "bab", "abc", "bca", "cab", and "abc".
The good substrings are "abc", "bca", "cab", and "abc".
```

**Constraints:**

- `1 <= s.length <= 100`
- `s` consists of lowercase English letters.
*/
package main

func countGoodSubstrings(s string) int {
	var count = 0
	var fp, sp, tp string
	for i := 0; i < len(s)-3; i++ {
		fp = string(s[i])
		sp = string(s[i+1])
		tp = string(s[i+2])
		if fp != sp && sp != tp && fp != tp {
			count++
		}
	}
	return count
}

