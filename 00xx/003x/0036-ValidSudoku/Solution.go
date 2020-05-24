/*  Solution: Lists/Matrices of Hash Maps
    Time Complexity:    O(n^2)
    Space Complexity:   O(3*n^2)
 */

type RuleMap map[byte]bool

func isValidSudoku(board [][]byte) bool {
  rows := []RuleMap{}
  cols := []RuleMap{}
  boxes := [][]RuleMap{}

  for a := 0; a < 3; a++ {
    rowOfBoxes := []RuleMap{}
    for b := 0; b < 3; b++ {
      rows = append(rows, RuleMap{
        '1': false,
        '2': false,
        '3': false,
        '4': false,
        '5': false,
        '6': false,
        '7': false,
        '8': false,
        '9': false,
      })
      cols = append(cols, RuleMap{
        '1': false,
        '2': false,
        '3': false,
        '4': false,
        '5': false,
        '6': false,
        '7': false,
        '8': false,
        '9': false,
      })
      rowOfBoxes = append(rowOfBoxes, RuleMap{
        '1': false,
        '2': false,
        '3': false,
        '4': false,
        '5': false,
        '6': false,
        '7': false,
        '8': false,
        '9': false,
      })
    }

    boxes = append(boxes, rowOfBoxes)
  }

  for i, row := range board {
    for j, cell := range row {
      if cell == '.' {
        continue
      }
      if rows[i][cell] == true {
        return false
      } else {
        rows[i][cell] = true
      }
      if cols[j][cell] == true {
        return false
      } else {
        cols[j][cell] = true
      }
      if boxes[i/3][j/3][cell] == true {
        return false
      } else {
        boxes[i/3][j/3][cell] = true
      }
    }
  }
  return true
}
