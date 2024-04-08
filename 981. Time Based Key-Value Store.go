type Pair struct {
    val string
    timeStamp int
}

type TimeMap struct {
    timeMap map[string][]Pair    
}


func Constructor() TimeMap {
    return TimeMap {
        timeMap: make(map[string][]Pair),
    }    
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    pair := Pair {
        val: value,
        timeStamp: timestamp,
    }
    if _, ok := this.timeMap[key]; !ok {
        this.timeMap[key] = make([]Pair, 0)
    }    
    this.timeMap[key] = append(this.timeMap[key], pair)
}


func (this *TimeMap) Get(key string, timestamp int) string {
    if _, ok := this.timeMap[key]; !ok {
        return ""
    } 
    arr := this.timeMap[key]
    if arr[0].timeStamp > timestamp {
        return ""
    } 
    left := 0
    right := len(arr)-1
    for left < right{
        mid := int((left+right)/2)
        if arr[mid].timeStamp == timestamp {
            return arr[mid].val
        } else if arr[mid].timeStamp > timestamp {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    if arr[left].timeStamp > timestamp {
        return arr[left-1].val
    }
    return arr[left].val
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */