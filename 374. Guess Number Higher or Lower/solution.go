/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

// Attempt 1: Recursive Solution 

// Does not work as a fatal runtime error occurs due to lack of memory
// func guessNumber(n int) int {
//     if guess(n) == 0 {
//         return n
//     } else if guess(n) < 0 {
//         return guessNumber(n - 1)
//     } else {
//         return guessNumber(n + 1)
//     }
// }

// Attempt 2: With Binary Search
// Main Idea: Repeatedly halving the search interval 
// Time Complexity: O(log(n))
// Space Complexity: O(1)

func guessNumber(n int) int {
    // Declare and initialize `left`, `right` and `middle` to 1, n and 0 respectively.
    left, right, middle := 1, n, 0

    // While left is less than or equal to right
    for left <= right {
        // `>>` means divided by 2
        // (left + right) divided by 2, 1 time
        middle = (left + right) >> 1

        // Declare `guess` and assign it the value of using the `guess` function on `middle`.
        guess := guess(middle)
        // If `guess` == -1, the number must be smaller than `middle` but larger than `left`.
        if guess == -1 {
            right = middle - 1
        
        // If `guess` == 1, the number must be bigger than `middle` but smaller than `right`.
        } else if guess == 1 {
            left = middle + 1
        
        // Else, `guess` == 0 and the number must be `n`.
        } else {
            break
        }
    }

    // Return `middle` which should be the number picked.
    return middle
}