/*  Solution: Stack
    Time Complexity:    O(n)
    Space Complexity:   O(n)
 */

func isValid(s string) bool {
  stack := []rune{}
  for _, char := range s {
    switch char {
    case '(':
      stack = append(stack, char)
    case '[':
      stack = append(stack, char)
    case '{':
      stack = append(stack, char)
    case '}':
      if len(stack) == 0 || stack[len(stack)-1] != '{' {
        return false
      }
      stack = stack[:len(stack)-1]
    case ']':
      if len(stack) == 0 || stack[len(stack)-1] != '[' {
        return false
      }
      stack = stack[:len(stack)-1]
    case ')':
      if len(stack) == 0 || stack[len(stack)-1] != '(' {
        return false
      }
      stack = stack[:len(stack)-1]
    }
  }
  if len(stack) > 0 {
    return false
  }
  return true
}
