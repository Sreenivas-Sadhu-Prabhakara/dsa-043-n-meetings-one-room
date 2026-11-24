package main

import (
	"fmt"
	"sort"
)

// MaxMeetings returns how many meetings fit in one room when a new meeting
// may start only strictly after the previous one ends. Greedy on earliest
// end time.
func MaxMeetings(start, end []int) int {
	n := len(start)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool { return end[idx[a]] < end[idx[b]] })
	count, lastEnd := 0, -1
	for _, i := range idx {
		if start[i] > lastEnd {
			count++
			lastEnd = end[i]
		}
	}
	return count
}

func main() {
	fmt.Println(MaxMeetings([]int{1, 3, 0, 5, 8, 5}, []int{2, 4, 6, 7, 9, 9}))
}
