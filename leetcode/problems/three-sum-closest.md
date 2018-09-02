# [16. 最接近的三数之和](https://leetcode-cn.com/problems/3sum-closest/description/)

## 第一次提交 200ms 12.16%

![提交记录](./three-sum-closest.png)

## 范例 8ms

```golang
func threeSumClosest(nums []int, ttarget int) int {
    sort.Ints(nums)
    
    
    l := len(nums)
    if l <= 2 {
        return -1
    }
    
    r := nums[0] + nums[1] + nums[2]
    
    m := nums[0]-1
    for k := 0; k < l-2; k++ {
        if m == nums[k] {
            continue
        }
        m = nums[k]
        target := ttarget - nums[k]
        
        i := k + 1
        j := l - 1
        for ; i < j ; {
            t := nums[i] + nums[j] + m
            if abs(t - ttarget) < abs(r - ttarget) {
                r = t
            }
            
            if nums[i] + nums[j] < target {
                i++
            } else if nums[i] + nums[j] > target {
                j--
            } else {
                return ttarget
            }
        }
    }
    
    return r
}


func abs( a int ) int {
    if a > 0 {
        return a
    }
    
    return 0 - a
}
```
