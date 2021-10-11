// 346. Moving Average from Data Stream
// https://leetcode.com/problems/moving-average-from-data-stream/

package main

import "fmt"

type MovingAverage struct {
	Size int
	Nums []int
	Sum  int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		Size: size,
		Nums: []int{},
		Sum:  0,
	}
}

// Golang doesn't jhave native queueing so whatever
func (this *MovingAverage) Next(val int) float64 {
	if len(this.Nums) == this.Size {
		this.Sum -= this.Nums[0]
		this.Nums = this.Nums[1:]
	}
	this.Sum += val
	this.Nums = append(this.Nums, val)
	return float64(this.Sum) / float64(len(this.Nums))
}

func main() {
	obj := Constructor(3)
	fmt.Println(obj.Next(1))
	fmt.Println(obj.Next(10))
	fmt.Println(obj.Next(3))
	fmt.Println(obj.Next(5))
}
