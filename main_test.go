package main

import (
	"testing"
)

func TestInitMinQueue(t *testing.T) {
	initVals := []int{113, 12, 1, 11, 111, 112, 114}

	heapRef := NewMinBinHeap(initVals)

	heapSlice := *heapRef

	t.Logf("\nheap: %v", heapSlice)

	prevNum := 1111

	for _, heapNum := range heapSlice {
		if heapNum > prevNum {
			//t.Errorf("\n prev %v cur %v", prevNum, heapNum)
		}

		prevNum = heapNum
	}
}

func TestPoll(t *testing.T) {
	initVals := []int{113, 12, 1, 11, 111, 112, 114}

	heapRef := NewMinBinHeap(initVals)

	heapRef.Poll()

	heapSlice := *heapRef

	t.Logf("\nheap: %v", heapSlice)

	for i, heapNum := range heapSlice {
		if heapNum == 114 && i < 3 {
			t.Errorf("\n largest val wrong depth")
		}
	}
}

func TestRemove(t *testing.T) {
	t.Run("Remove as poll", func(t *testing.T) {
		initVals := []int{113, 12, 1, 11, 111, 112, 114}

		heapRef := NewMinBinHeap(initVals)

		heapRef.Remove(1)

		heapSlice := *heapRef

		t.Logf("\nheap: %v", heapSlice)

		for i, heapNum := range heapSlice {
			if heapNum == 114 && i < 3 {
				t.Errorf("\n largest val wrong depth")
			}
		}
	})
	t.Run("Remove for bubble down", func(t *testing.T) {
		initVals := []int{113, 12, 1, 11, 111, 112, 114}

		heapRef := NewMinBinHeap(initVals)

		heapRef.Remove(11)

		heapSlice := *heapRef

		t.Logf("\nheap: %v", heapSlice)

		for i, heapNum := range heapSlice {
			if heapNum == 114 && i < 3 {
				t.Errorf("\n largest val wrong depth")
			}
		}
	})
}
