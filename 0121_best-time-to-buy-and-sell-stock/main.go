// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package main

import "fmt"

func maxProfit_brute_force(prices []int) int {
	// This will violate the time limit
	p := 0
	for i := 0; i < len(prices); i++ {
		for j := len(prices) - 1; j > i; j-- {
			if prices[j]-prices[i] > p {
				p = prices[j] - prices[i]
			}
		}
	}
	return p
}

func maxProfit(prices []int) int {
	max := 0
	p := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		}
		if max-prices[i] > p {
			p = max - prices[i]
		}
	}
	return p
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
