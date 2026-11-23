class Solution:
    def isValid(self, s: str) -> bool:
        stk = []
        for c in s:
            if c == '[' or c =='(' or c =='{':
                stk.append(c)
                continue
            if len(stk) == 0:
                return False
            v = stk.pop()
            if c == ']':
                if v != '[':
                    return False
            if c == '}':
                if v != '{':
                    return False
            if c == ')':
                if v != '(':
                    return False
        if len(stk) > 0:
            return False
        return True