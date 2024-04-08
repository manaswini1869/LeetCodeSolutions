class MinStack:

    def __init__(self):
        self.res = [] 
        self.minres = []       

    def push(self, val: int) -> None:
        self.res.append(val)
        if self.minres:
            val = min(val, self.minres[-1])
        self.minres.append(val)
        

    def pop(self) -> None:
        self.res.pop()
        self.minres.pop()
        

    def top(self) -> int:
        return self.res[-1]
        

    def getMin(self) -> int:
        return self.minres[-1]
        
        
        


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()