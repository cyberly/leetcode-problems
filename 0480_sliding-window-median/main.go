// 480. Sliding Window Median
// https://leetcode.com/problems/sliding-window-median/

package main

import (
	"container/list"
	"fmt"
	"sort"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	ans := []float64{}
	// Create our initial window list
	// Add median to the result
	win := make([]int, k)
	copy(win, nums)
	sort.Ints(win)
	fmt.Println(win)
	l := list.New()
	for i := 0; i < len(win); i++ {
		l.PushBack(win[i])
	}
	ans = append(ans, getMedian(l, k))
	for i := k; i < len(nums); i++ {
		l = removeListValue(l, nums[i-k])
		l = inserttListValue(l, nums[i])
		ans = append(ans, getMedian(l, k))
	}
	return ans
}

func getMedian(l *list.List, k int) float64 {
	list := l.Front()
	for i := 0; i < k/2; list, i = list.Next(), i+1 {
	}
	if k%2 == 0 {
		prev := list.Prev()
		return (float64(prev.Value.(int)) + float64(list.Value.(int))) / 2
	}
	return float64(list.Value.(int))
}

func inserttListValue(l *list.List, n int) *list.List {
	for i := l.Front(); i != nil; i = i.Next() {
		if i.Value.(int) >= n {
			l.InsertBefore(n, i)
			return l
		}
	}
	l.PushBack(n)
	return l
}

func removeListValue(l *list.List, n int) *list.List {
	for i := l.Front(); i != nil; i = i.Next() {
		if i.Value.(int) == n {
			l.Remove(i)
		}
	}
	return l
}

func medianSlidingWindow_slow(nums []int, k int) []float64 {
	// Pretty sure this works fine but overdlows timelimit cause it is bad
	var med float64
	ans := []float64{}
	if k == 1 {
		for i := 0; i < len(nums); i++ {
			ans = append(ans, float64(nums[i]))
		}
		return ans
	}
	window := make([]int, k)
	topEnd := len(nums) - k
	sort.Ints(window)
	for i := 0; i <= topEnd; i++ {
		//fmt.Printf("I: %v, k: %v\n", i, k)
		copy(window, nums[0:k])
		sort.Ints(window)
		//fmt.Println(window)
		if k%2 == 0 {
			med = (float64(window[k/2]) + float64(window[k/2-1])) / 2
		} else {
			med = float64(window[k/2])
		}
		ans = append(ans, med)
		nums = nums[1:]
	}
	return ans
}

// This one was from the discussion and I don't fully understand it so leaving it here to circle back
// It's like 50-60ms and way faster than my version above
func medianSlidingWindow_bin(nums []int, k int) []float64 {
	ans := []float64{}
	if k == 1 {
		for i := range nums {
			ans = append(ans, float64(nums[i]))
		}
		return ans
	}

	window := make([]int, k)
	copy(window, nums[0:k])
	sort.Ints(window)
	for i := 0; i+k <= len(nums); i++ {
		ans = append(ans, getMedian_bin(window))
		idx := findTarget(window, nums[i])
		if i+k >= len(nums) {
			break
		}

		window[idx] = nums[i+k]
		for idx > 0 && window[idx] < window[idx-1] {
			window[idx], window[idx-1] = window[idx-1], window[idx]
			idx--
		}
		for idx+1 < k && window[idx] > window[idx+1] {
			window[idx], window[idx+1] = window[idx+1], window[idx]
			idx++
		}
	}

	return ans
}

func findTarget(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for start+1 < end {
		mid := start + (end-start)/2

		if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if nums[start] == target {
		return start
	}
	return end
}

func getMedian_bin(nums []int) float64 {
	arrLen := len(nums)

	if arrLen%2 == 0 {
		return (float64(nums[arrLen/2]+nums[arrLen/2-1]) / 2.0)
	}
	return float64(nums[arrLen/2])
}

func main() {
	//fmt.Println(medianSlidingWindow([]int{1, 2, 3, 4, 2, 3, 1, 4, 2}, 3))
	fmt.Println(medianSlidingWindow([]int{1, 4, 2, 3}, 4))

}
