func isValid(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    
    for _, c := range s {
        // If it's an opening bracket, push to stack
        if c == '(' || c == '[' || c == '{' {
            stack = append(stack, c)
            continue
        }
        
        // If it's a closing bracket
        if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
            return false
        }
        stack = stack[:len(stack)-1]
    }
    
    return len(stack) == 0
}