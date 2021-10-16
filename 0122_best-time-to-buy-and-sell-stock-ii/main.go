// 122. Best Time to Buy and Sell Stock II
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

package main

import "fmt"

func maxProfit(prices []int) int {
	maxProf := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProf += prices[i] - prices[i-1]
		}
	}
	return maxProf
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
