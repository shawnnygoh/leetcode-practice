func isValid(s string) bool {
    // Base case: If length of string s is equal to 0 or is odd (not pairs)
    if len(s) == 0 || len(s) % 2 == 1 {
        return false
    } else {
        // Create a map of runes to store parentheses pairings
        parentheses := map[rune]rune {
            '(': ')',
            '{': '}',
            '[': ']',
        }
        
        // Create a stack data structure (array to store opening parenthesis while iterating through string)
        stack := []rune{}

        // Loop through the string and...
        for _, i := range s {

            // If it is an opening parenthesis 
            if _, ok := parentheses[i]; ok {

                // Push opening parenthesis into stack
                stack = append(stack, i)   

            // If stack is empty or top element in stack is not the corresponding opening parenthesis for the 
            // current element (which should be a closing parenthesis)
            } else if len(stack) == 0 || parentheses[stack[len(stack) - 1]] != i {
                return false
            
            // If the opening and closing parentheses match, remove the top element from the stack
            } else {
                stack = stack[:len(stack) - 1]
            }
        }

        // If length of stack is empty at the end, all opening parentheses have been matched with closing 
        // parentheses and the string is valid
        return len(stack) == 0
    }
}