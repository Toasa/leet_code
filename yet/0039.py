class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        memo = [0] * len(candidates)
        for item in candidates:
            memo[item] = 1

        l = []
        for item in candidates:
            count = 1
            for item * count <= target:
                if memo[target - item * count] == 1:
                    tmp = [item] * count
                    tmp.append(target - item * count)
                    l.append(tmp)
                count += 1


            if target % item == 0:
                # １つの要素からなるlistの作成
                l.append([item] * (target // item))

        return l