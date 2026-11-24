package main

import "testing"

func TestMaxMeetings(t *testing.T) {
	cases := []struct {
		start, end []int
		want       int
	}{
		{[]int{1, 3, 0, 5, 8, 5}, []int{2, 4, 6, 7, 9, 9}, 4},
		{[]int{10, 12, 20}, []int{20, 25, 30}, 1}, // 12 starts before 20 ends; 20 does not start AFTER 20
		{[]int{1}, []int{2}, 1},
		{[]int{}, []int{}, 0},
		{[]int{1, 2}, []int{10, 11}, 1},
	}
	for _, c := range cases {
		if got := MaxMeetings(c.start, c.end); got != c.want {
			t.Errorf("MaxMeetings(%v,%v) = %d, want %d", c.start, c.end, got, c.want)
		}
	}
}
