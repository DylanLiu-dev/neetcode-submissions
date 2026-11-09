class Solution:
    def isColumnDup(self, board, i, j):
        seen = set()
        for c in range(0, 9):
            if board[c][j] == ".":
                continue
            if board[c][j] in seen:
                return True
            seen.add(board[c][j])
        return False
    def isRowDup(self, board, i, j):
        seen = set()
        for r in range(0, 9):
            if board[i][r] == ".":
                continue
            if board[i][r] in seen:
                return True
            seen.add(board[i][r])
        return False
    def isBlockDup(self, board, i, j):
        toplefti = i//3 * 3
        topleftj = j//3 * 3
        seen = set()
        for c in range(3):
            for r in range(3):
                v = board[toplefti+c][topleftj+r]
                if v == ".":
                    continue
                if v in seen:
                    return True
                seen.add(v)
        return False
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for i in range(9):
            for j in range(9):
                if self.isColumnDup(board, i, j) \
                or self.isRowDup(board, i, j) \
                or self.isBlockDup(board, i, j):
                    return False
        return True
