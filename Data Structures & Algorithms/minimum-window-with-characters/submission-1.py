class Solution:
    def minWindow(self, s: str, t: str) -> str:
        count = {}
        for c in t:
            count[c] = 1 + count.get(c, 0)
        l = 0
        have, need = 0, len(count)
        res = ""
        window = {}
        for r in range(len(s)):
            c = s[r]
            window[c] = 1 + window.get(c, 0)
            if c in count and window[c] == count[c]:
                have += 1
            while have == need:
                if res=="" or r-l+1 < len(res):
                    res = s[l:r+1]
                window[s[l]] -= 1
                if s[l] in count and window[s[l]] < count[s[l]]:
                    have -= 1
                l += 1
        return res