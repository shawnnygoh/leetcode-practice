// Main Idea: Fibonacci Sequence

// Attempt 1: Recursive Solution
// Time Complexity: O(2^n)
// Space Complexity: O(n)

// Does not work, results in Time Limit Exceeded as previously calculated values are 
// recalculated.
// func climbStairs(n int) int {
//     if n < 3 {
//         return n
//     } else {
//         return climbStairs(n - 1) + climbStairs(n - 2)
//     }
// }


// Attempt 2: Iterative Solution
// Time Complexity: O(n)
// Space Complexity: O(1)
func climbStairs(n int) int {
    distinct := 0

    n1 := 0
    n2 := 0

    for i := 1; i <= n; i++ {
        if i < 3 {
            distinct = i
        } else {
            distinct = n1 + n2
        }

        n2 = n1 
        n1 = distinct 
    }

    return distinct 
}