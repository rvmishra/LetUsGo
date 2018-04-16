package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(bsearch(a, 2))
}

func bsearch(a []int, k int) int {
	if a == nil || len(a) == 0 {
		return -1
	}
	l := 0
	r := len(a) - 1
	for l <= r {
		mid := l + (r-l)/2
		if a[mid] == k {
			return mid
		} else if k < a[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
