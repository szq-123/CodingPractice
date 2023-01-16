package questions

import (
	"fmt"
	"sort"
	"testing"
)

func TestCountingSort(t *testing.T) {
	fmt.Printf("%v", CountingSort([]int{1, 3, 4, 3, 2, 1}, 5, func(num int) int {
		return num
	}))
}

func TestRadixSort(t *testing.T) {
	fmt.Printf("%v", RadixSort([]int{13, 32, 41, 33, 22, 11}, 10, 2))
}

func TestFindNumKByQuickSort(t *testing.T) {
	arr := []int{13, 32, 41, 33, 22, 11}
	fmt.Printf("%v\n", FindNumKByQuickSort(arr, 0, len(arr), 4))
	sort.Ints(arr)
	fmt.Printf("%v", arr)
}

func TestInsertSort(t *testing.T) {
	arr := []int{13, 32, 41, 33, 22, 11}
	InsertSort(arr, 0, len(arr))
	fmt.Printf("%v", arr)
}

func TestFindNumK(t *testing.T) {
	arr := []int{13, 32, 41, 33, 22, 11}
	fmt.Printf("%v\n", FindNumK(arr, 0, len(arr), 4))
	sort.Ints(arr)
	fmt.Printf("%v", arr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 2, 3, 4, 2, 1, 5, 3, 2, 3, 4, 3, 32, 4, 1, 2345, 2}
	QuickSort(arr)
	fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
	arr := []int{3, 2, 3, 4, 2, 1, 5, 3, 2, 3, 4, 3, 32, 4, 1, 2345, 2}
	HeapSort(arr)
	fmt.Println(arr)
}

func TestPickPrimeN(t *testing.T) {
	fmt.Println(PickPrimeN3(10))
}
