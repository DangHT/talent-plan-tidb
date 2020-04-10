package main

import (
	"fmt"
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.

var wg sync.WaitGroup

func main() {
	arr := []int64{2, 3, 1, 5, 9, 6, 4, 8, 7}
	MergeSort(arr)
	for i := range arr {
		fmt.Printf("%d ", i)
	}
}

// Test results:
// 	make test: 1.2s ~ 1.3s
// 	make bench: 46s ~ 47s
func MergeSort(src []int64) {
	mergeSort(src, len(src))
}

func mergeSort(arr []int64, n int) {
	var length int = 1
	tmpArr := make([]int64, len(arr))
	for length < n {
		mergePass(arr, tmpArr, n, length)
		length *= 2
		mergePass(tmpArr, arr, n, length)
		length *= 2
	}
}

func mergePass(arr, tmpArr []int64, n, length int) {
	i := 0
	for ; i <= n-2*length; i += 2 * length {
		wg.Add(1)
		go merge(arr, tmpArr, i, i+length, i+2*length-1)
	}
	wg.Wait()
	if i+length < n {
		merge(arr, tmpArr, i, i+length, n-1)
	} else {
		for j := i; j < n; j++ {
			tmpArr[j] = arr[j]
		}
	}
}

func merge(arr, tmpArr []int64, l, r, re int) {
	le := r - 1
	tmp := l
	for l <= le && r <= re {
		if arr[l] <= arr[r] {
			tmpArr[tmp] = arr[l]
			l++
		} else {
			tmpArr[tmp] = arr[r]
			r++
		}
		tmp++
	}
	for l <= le {
		tmpArr[tmp] = arr[l]
		tmp++
		l++
	}
	for r <= re {
		tmpArr[tmp] = arr[r]
		tmp++
		r++
	}
}
