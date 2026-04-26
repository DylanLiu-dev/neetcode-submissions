func isValidSudoku(board [][]byte) bool {
	for _, b := range board {
		if hasDuplicate(b) {
			return false
		}
	}
	for i:=0; i<9; i++ {
		tmp := make([]byte, 0, 9)
		for j:=0; j<9; j++ {
			tmp = append(tmp, board[j][i])
		}
		if hasDuplicate(tmp) {
			return false
		}
	}
	for x:=0; x<3; x++ {
		for y:=0; y<3; y++ {
			tmp := make([]byte, 0, 9)
			for j:=0; j<3; j++ {
				for k:=0; k<3; k++ {
					tmp = append(tmp, board[x*3+j][y*3+k])
				}
			}
			if hasDuplicate(tmp) {
				return false
			}
		}
	}
	return true
}

func hasDuplicate(nums []byte) bool {
	seen := map[byte]struct{}{}
	for _, num := range nums {
		if num == '.' {
			continue
		}
		if _, exist := seen[num]; exist {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}