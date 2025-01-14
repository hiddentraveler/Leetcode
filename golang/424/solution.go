// Source: https://leetcode.com/problems/longest-repeating-character-replacement/description/
// Author: Neo Orez
// Difficulty: Medium

/*
You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return *the length of the longest substring containing the same letter you can get after performing the above operations*.

**Example 1:**

```
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
```

**Example 2:**

```
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.
```

**Constraints:**

- `1 <= s.length <= 105`
- `s` consists of only uppercase English letters.
- `0 <= k <= s.length`
*/
package main

func characterReplacement(s string, k int) int {

	var chain = 0
	var hash = make(map[string]int)
	var i = 0
	for b := 0; b < len(s); b++ {
		letter := s[b]
		hash[string(letter)]++
		mF := maxFreq(hash)
		currentLength := b - i + 1
		if currentLength-mF > k {
			hash[string(s[i])]--
			i++
			currentLength = b - i + 1
		}
		if currentLength > chain {
			chain = currentLength
		}
	}
	return chain
}

func maxFreq(m map[string]int) int {
	max := 0
	for _, value := range m {
		if value > max {
			max = value
		}
	}
	return max
}

