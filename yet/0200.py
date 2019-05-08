class Solution:
    land = []
    sea = []
    island_count = 0
    visit_count = 0

    # checkedが0の場合、3の適切な処理をする
    def assort(self, check_val, grid_val, r, c):
        if check_val == 0:
            # 海の場合
            if grid_val == 0:
                Solution.sea.append([r, c])
            # 陸の場合
            else:
                Solution.land.append([r, c])

    def numIslands(self, grid: List[List[str]]) -> int:
        row = len(grid)
        col = len(grid[0])
        checked = [[0 for i in range(col)] for j in range(row)]

        # 1. 最初にgrid[0][0]を検査し, landかseaに入れる(enqueする)
        #
        # 2-1. landが空でない場合dequeし3へ
        # 2-2. landが空の場合
        #   2-2-1. すべてのcellに訪れていた場合 -> return island_count　する
        #   2-2-2. island_count++し、sea[]からdequeし、3へ
        #
        # 3. visit_count++する。checkedにチェックする。
        #   取り出したcellに隣接するcellを調べる
        #   調べ方、checkedが1の場合、すでに訪れているので、何もしない
        #         checkedが0の場合  
        #                         海の場合 -> seaにenqueする
        #                         陸の場合 -> landにenqueする
        #
        # 4. 2に戻る

        # 1
        if grid[0][0] == "1":
            Solution.land.append([0, 0])
        else:
            Solution.sea.append([0, 0])

        # 手順2で取り出すcellの宣言
        cell = [0, 0]
        while True:

            if len(Solution.land) > 0:
                cell = Solution.land[0]
                Solution.land = Solution.land[1:]
            else:
                if Solution.visit_count == col * row:
                    return Solution.island_count
                elif len(Solution.sea) == 0:
                    return Solution.island_count
                else:
                    Solution.island_count += 1
                    cell = Solution.sea[0]
                    Solution.sea = Solution.sea[1:]

            # 3.
            Solution.visit_count += 1
            row_i = cell[0]
            col_i = cell[1]


            print("cell:", cell, "land:", Solution.land)
            for i in range(row):
                for j in range(col):
                    print(checked[i][j], end="")
                print()
            print("***********************")


            checked[row_i][col_i] = 1

            if row_i > 1:
                self.assort(checked[row_i - 1][col_i], grid[row_i - 1][col_i], row_i - 1, col_i)
            if row_i < row - 1:
                self.assort(checked[row_i + 1][col_i], grid[row_i + 1][col_i], row_i + 1, col_i)
            if col_i > 1:
                self.assort(checked[row_i][col_i - 1], grid[row_i][col_i - 1], row_i, col_i - 1)
            if col_i < col - 1:
                self.assort(checked[row_i][col_i + 1], grid[row_i][col_i + 1], row_i, col_i + 1)