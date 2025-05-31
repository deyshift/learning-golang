package main

func lengthOfLongestSubstring(s string) int {
    start := 0
    maxLength := 0
    charIndexMap := make(map[rune]int)

    for end, char := range s {
        if index, ok := charIndexMap[char]; ok && index >= start {
            start = index + 1
        }
        charIndexMap[char] = end
        currentLength := end - start + 1
        if currentLength > maxLength {
            maxLength = currentLength
        }
    }

    return maxLength
}

func main() {
    // Example Usage
    println(lengthOfLongestSubstring("abcabcbb"))   // Output: 3
    println(lengthOfLongestSubstring("bbbbb"))      // Output: 1
    println(lengthOfLongestSubstring("pwwkew"))     // Output: 3
    println(lengthOfLongestSubstring(""))           // Output: 0
    println(lengthOfLongestSubstring("dvdf"))       // Output: 3
    println(lengthOfLongestSubstring("anviaj"))     // Output: 5
    println(lengthOfLongestSubstring("tmmzuxt"))    // Output: 5
}