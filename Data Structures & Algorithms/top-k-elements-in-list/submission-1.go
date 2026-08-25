func topKFrequent(nums []int, k int) []int {
    // 1. Count frequency of each number
    freq := make(map[int]int)

    for _, num := range nums {
        freq[num]++
    }

    // 2. Create buckets
    // bucket[i] contains numbers that appear i times
    buckets := make([][]int, len(nums)+1)

    for num, count := range freq {
        buckets[count] = append(buckets[count], num)
    }

    // 3. Start from highest frequency
    result := []int{}

    for count := len(buckets) - 1; count >= 0 && len(result) < k; count-- {
        for _, num := range buckets[count] {
            result = append(result, num)

            if len(result) == k {
                return result
            }
        }
    }

    return result
}
