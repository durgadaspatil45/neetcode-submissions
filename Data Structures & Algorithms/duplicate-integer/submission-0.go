func hasDuplicate(nums []int) bool {
    mp := make(map[int]struct{})
    for _, key := range nums{
        if _, ok := mp[key]; ok {
            return true
        }
        mp[key] = struct{}{}
    }
    return false
}