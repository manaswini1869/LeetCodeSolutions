class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for i in range(9):
            k = set()
            l = set()
            for j in range(9):
                if board[i][j] in k or board[j][i] in l:
                    return False
                if board[i][j] != '.':
                    k.add(board[i][j])
                if board[j][i] != '.':
                    l.add(board[j][i])
        for i in range(0, 9, 3):
            for j in range(0, 9, 3):
                subgrid_set = set()
                for k in range(3):
                    for l in range(3):
                        if board[i + k][j + l] in subgrid_set:
                            return False
                        if board[i + k][j + l] != '.':
                            subgrid_set.add(board[i + k][j + l])
        return True
        
        

        