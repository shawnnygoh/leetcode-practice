// Main Idea: Find the GCD of the lengths of `str1` and `str2` and use it to slice `str1` to obtain the GCD of the strings.
// Time Complexity: O(m + n)
// Space Complexity: O(m + n)

func gcdOfStrings(str1 string, str2 string) string {
    // If a GCD is not possible, return a blank string.
    if str1 + str2 != str2 + str1 {
        return ""
    }

    // Else, initialize and declare `slicelength` as the GCD of the lengths of `str1` and `str2`.
    slicelength := gcd(len(str1), len(str2))

    // Return the GCD of `str1` and `str2` by slicing the `slicelength` from `str1`.
    return str1[:slicelength]
}

// Function to calculate the greatest common divisor of two integers.
func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}