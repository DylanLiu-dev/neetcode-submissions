class Solution:
    def minWindow(self, s: str, t: str) -> str:
        count = {}
        for c in t:
            count[c] = 1 + count.get(c, 0)
        l = 0
        have, need = 0, len(count)
        res_l, res_len = 0, float("inf")
        window = {}
        for r in range(len(s)):
            c = s[r]
            window[c] = 1 + window.get(c, 0)
            if c in count and window[c] == count[c]:
                have += 1
            while have == need:
                if r-l+1 < res_len:
                    res_l = l
                    res_len = r-l+1
                window[s[l]] -= 1
                if s[l] in count and window[s[l]] < count[s[l]]:
                    have -= 1
                l += 1
        return s[res_l:res_l+res_len] if res_len != float("inf") else ""