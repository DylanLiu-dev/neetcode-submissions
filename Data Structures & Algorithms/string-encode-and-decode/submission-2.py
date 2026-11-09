class Solution:

    def encode(self, strs: List[str]) -> str:
        res = ""
        for string in strs:
            res += str(len(string))
            res += "#"
            res += string
        return res
    def decode(self, s: str) -> List[str]:
        res = []
        i = 0
        while i < len(s):
            tmp = ""
            while s[i]!="#":
                tmp += s[i]
                i+=1
            length = int(tmp)
            res.append(s[i+1:i+length+1])
            i += length+1
        return res