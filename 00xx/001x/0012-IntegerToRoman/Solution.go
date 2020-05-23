/*  Solution: Hash Map
    Time Complexity:    O(nk) (n=3, check all k keys of map 3 times)
    Space Complexity:   O(n+2k) (map and int slice of size k + return str)
 */

func intToRoman(num int) string {
  roman := ""
  table := map[int]string{
    1000: "M",
    900: "CM",
    500: "D",
    400: "CD",
    100: "C",
    90: "XC",
    50: "L",
    40: "XL",
    10: "X",
    9: "IX",
    5: "V",
    4: "IV",
    1: "I",
  }
  //  In Go, loops iterate through map keys in RANDOM order
  //  To bypass this, have a separated ordered list of keys to iterate
  keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
  for num > 0 {
    for _, key := range keys {
      if num >= key {
        roman += table[key]
        num -= key
        break
      }
    }
  }
  return roman
}
