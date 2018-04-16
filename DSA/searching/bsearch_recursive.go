package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(bsearch_rec(a, 6, 0, len(a)-1))
}

func bsearch_rec(a []int, k int, l int, r int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if a[mid] == k {
		return mid
	} else if k < a[mid] {
		return bsearch_rec(a, k, l, mid-1)
	} else {
		return bsearch_rec(a, k, mid+1, r)
	}
}
