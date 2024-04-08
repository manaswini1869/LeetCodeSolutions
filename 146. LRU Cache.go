type Node struct {
    key int
    val int
    prev *Node
    next *Node
}

type LRUCache struct { 
    cap int
    cache map[int]*Node
    left *Node
    right *Node
}


func Constructor(capacity int) LRUCache {
    leftFake := &Node{val: -1, prev: nil, next: nil}
    rightFake := &Node{val: -1, prev: leftFake, next: nil}
    leftFake.next = rightFake
    return LRUCache{
        cap: capacity,
        cache: make(map[int]*Node, capacity),
        left: leftFake,
        right: rightFake,
    }
    
}

func (this *LRUCache) remove(node *Node) {
    prev := node.prev
    nxt := node.next
    prev.next = nxt
    nxt.prev = prev
}

func (this *LRUCache) insert(node *Node) {
    prev := this.right.prev
    nxt := this.right
    prev.next = node
    nxt.prev = node
    node.next = nxt 
    node.prev = prev
}

func (this *LRUCache) Get(key int) int {
    if _,ok:= this.cache[key]; !ok {
        return -1
    }  
    this.remove(this.cache[key])
    this.insert(this.cache[key])
    return this.cache[key].val
}


func (this *LRUCache) Put(key int, value int)  {
    if _,ok := this.cache[key]; ok {
        this.remove(this.cache[key])
    }
    node := &Node{
        key: key,
        val: value,
    }
    this.cache[key] = node
    this.insert(this.cache[key])

    if len(this.cache) > this.cap {
        lru := this.left.next
        this.remove(lru) // delete from the linked list
        delete(this.cache, lru.key) // delete the lru from the map
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */