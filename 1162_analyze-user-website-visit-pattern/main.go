// 1152. Analyze User Website Visit Pattern
// https://leetcode.com/problems/analyze-user-website-visit-pattern/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	s := make([]int, len(username))
	for i := range s {
		s[i] = i
	}
	sort.Slice(s, func(i, j int) bool {
		// Thanks, StackOverflow
		if username[s[i]] == username[s[j]] {
			return timestamp[s[i]] < timestamp[s[j]]
		} else {
			return username[s[i]] < username[s[j]]
		}
	})
	siteSeq := make(map[string]int)
	userSeq := make(map[string]string)
	var maxSeq string
	maxSeqCount := 0
	for i := 0; i < len(username); i++ {
		for j := i + 1; j < len(username) && username[s[i]] == username[s[j]]; j++ {
			for k := j + 1; k < len(username) && username[s[i]] == username[s[k]]; k++ {
				seq := website[s[i]] + "," + website[s[j]] + "," + website[s[k]]
				if userSeq[seq] == username[s[i]] {
					continue
				} else {
					userSeq[seq] = username[s[i]]
				}
				siteSeq[seq] += 1
				if siteSeq[seq] > maxSeqCount {
					maxSeqCount = siteSeq[seq]
					maxSeq = seq
				} else {
					if siteSeq[seq] == maxSeqCount {
						if seq < maxSeq {
							maxSeq = seq
						}
					}
				}
			}
		}

	}
	return strings.Split(maxSeq, ",")
}

func main() {
	usernames := []string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"}
	websites := []string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"}
	timestamps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	usernames2 := []string{"dowg", "dowg", "dowg"}
	websites2 := []string{"y", "loedo", "y"}
	timestamps2 := []int{158931262, 562600350, 148438945}
	usernames3 := []string{"zkiikgv", "zkiikgv", "zkiikgv", "zkiikgv"}
	websites3 := []string{"wnaaxbfhxp", "mryxsjc", "oz", "wlarkzzqht"}
	timestamps3 := []int{436363475, 710406388, 386655081, 797150921}
	fmt.Println(mostVisitedPattern(usernames, timestamps, websites))
	fmt.Println(mostVisitedPattern(usernames2, timestamps2, websites2))
	fmt.Println(mostVisitedPattern(usernames3, timestamps3, websites3))
}
