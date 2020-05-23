/*  Solution: Two Hash Maps
    Time Complexity:    O(n)
    Space Complexity:   O(n+1) (Carry)

    Note: Is there a better solution...?
 */

func addStrings(num1 string, num2 string) string {
  atoi := map[byte]int{
    '0': 0,
    '1': 1,
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
  }

  itoa := map[int]string{
    0: "0",
    1: "1",
    2: "2",
    3: "3",
    4: "4",
    5: "5",
    6: "6",
    7: "7",
    8: "8",
    9: "9",
  }

  i := len(num1) - 1
  j := len(num2) - 1
  s := ""
  carry := 0

  for i > -1 || j > -1 {
    op1 := 0
    op2 := 0
    if i > -1 {
      if a, ok := atoi[num1[i]]; ok {
        op1 = a
        i--
      }
    }
    if j > -1 {
      if b, ok := atoi[num2[j]]; ok {
        op2 = b
        j--
      }
    }
    s = itoa[(op1 + op2 + carry) % 10] + s
    if op1 + op2 + carry > 9 {
      carry = 1
    } else {
      carry = 0
    }
  }
  if carry == 1 {
    return "1" + s
  } else {
    return s
  }
}
