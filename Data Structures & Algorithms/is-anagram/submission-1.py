class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        cnts = Counter(s)
        for c in t:
            cnts[c]-=1
        for k in cnts.keys():
            if cnts[k] != 0:
                return False
        return True