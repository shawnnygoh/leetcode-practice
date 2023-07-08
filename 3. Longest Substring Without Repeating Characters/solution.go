func max(firstnum, secondnum int) int {
    if firstnum > secondnum {
        return firstnum
    }
    return secondnum
}

func lengthOfLongestSubstring(s string) int {
    // Solution using two pointer sliding window:
    // Length of input string s 
    length := len(s)

    // Length of longest substring 
    var maxlength int

    // Map to store characters and their respective index
    charIndexMap := make(map[uint8]int) 

    // Start of current substring 
    var start int

    for end := 0; end < length; end++ {
        repeatIndex, isRepeated := charIndexMap[s[end]]
        if isRepeated {
            maxlength = max(maxlength, end - start)

            for i := start; i <= repeatIndex; i++ {
                delete(charIndexMap, s[i])
            }

            start = repeatIndex + 1
        }

        charIndexMap[s[end]] = end
    }

    maxlength = max(maxlength, length - start)
    return maxlength
}