// Main Idea: Traverse through the array and accumulate all the zeroes together by swapping elements around.
// Time Complexity: O(n), where n is the length of `nums`
// Space Complexity: O(1)

func moveZeroes(nums []int)  {

    // Declare and initialize a variable `zeroes` to keep track of the number of zeroes in `nums`.
    zeroes := 0

    // Iterate through `nums` with a `for` loop.
    for i := 0; i < len(nums); i++ {

        // If the number at index `i` is equal to 0, add 1 to the `zeroes` counter. 
        if nums[i] == 0 {
            zeroes++
        
        // If the number at index `i` is not equal to 0, and the `zeroes` counter is greater than 0, swap the 
        // leftmost 0 with the number at index `i`.
        } else if zeroes > 0 {
            nums[i - zeroes] = nums[i]
            nums[i] = 0
        }
    } 
}