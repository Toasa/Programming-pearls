package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := []int{}
	n := 10000000
	for len(s) != 1000000 {
		s = append(s, rand.Intn(n))
	}
	start := time.Now()
	// sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	arr := sort(s)
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	fmt.Println(arr)
}

func sort(s []int) [1000000]int {
	bits := [10000000]bool{}
	for _, val := range s {
		bits[val] = true
	}

	index := 0
	arr := [1000000]int{}
	for i := 0; i < 10000000; i++ {
		if bits[i] {
			arr[index] = i
			index++
		}
	}
	return arr
}
