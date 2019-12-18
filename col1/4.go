package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	solve()
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func solve() {
	// 0 < k < n
	const n int = 10000000
	const k int = 1000000
	list := [k]int{}
	picked := [n]bool{}

	var num int

	for index := 0; index < k; index++ {
		for true {
			num = rand.Intn(int(n))
			if !picked[num] {
				picked[num] = true
				break
			}
		}
		list[index] = num
	}
}
