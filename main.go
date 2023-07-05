package main

import (
	"fmt"
)

type MinBinHeap []int

func NewMinBinHeap(initNums []int) *MinBinHeap {
	minBinaryHeap := MinBinHeap{}
	heapRef := &minBinaryHeap

	for _, num := range initNums {
		heapRef.Insert(num)
	}

	return heapRef
}

func (minBinHeap *MinBinHeap) Remove(num int) {
	curHeap := *minBinHeap

	curIdx := -1

	for i, heapNum := range curHeap {
		if num == heapNum {
			curIdx = i
			break
		}
	}

	if curIdx == -1 {
		fmt.Errorf("\n num not in heap")
		return
	}

	updatedHeap := removeByIdx(curIdx, curHeap)

	*minBinHeap = updatedHeap
}

func (minBinHeap *MinBinHeap) Poll() {
	curHeap := *minBinHeap

	updatedHeap := removeByIdx(0, curHeap)

	*minBinHeap = updatedHeap
}

func removeByIdx(idx int, minHeap MinBinHeap) MinBinHeap {
	curIdx := idx
	fIdx := len(minHeap) - 1

	minHeap[curIdx], minHeap[fIdx] = minHeap[fIdx], minHeap[curIdx]

	fIdx--
	minHeap = minHeap[:fIdx]

	for {
		//bubble up
		// not needed for poll but oh well
		toEven := curIdx % 2

		parentIdx := ((curIdx + toEven) / 2) - 1

		parentVal := minHeap[0]

		if parentIdx > 0 {
			parentVal = minHeap[parentIdx]
		}

		if parentVal > minHeap[curIdx] {
			minHeap[parentIdx], minHeap[curIdx] = minHeap[curIdx], minHeap[parentIdx]
			curIdx = parentIdx
			continue
		}

		//bubble down
		if curIdx*2+2 < len(minHeap) {
			lChildIdx := (curIdx * 2) + 1
			rChildIdx := (curIdx * 2) + 2

			lChildVal := minHeap[lChildIdx]
			rChildVal := minHeap[rChildIdx]

			lesserIdx := lChildIdx

			if rChildVal < lChildVal {
				lesserIdx = rChildIdx
			}

			if minHeap[curIdx] > minHeap[lesserIdx] {
				minHeap[curIdx], minHeap[lesserIdx] = minHeap[lesserIdx], minHeap[curIdx]
				curIdx = lesserIdx
				continue
			}
		}

		break
	}

	return minHeap
}

func (minBinHeap *MinBinHeap) Insert(num int) {
	curHeap := *minBinHeap
	curHeap = append(curHeap, num)

	curIdx := len(curHeap) - 1

	for {
		toEven := curIdx % 2

		parentIdx := ((curIdx + toEven) / 2) - 1

		parentVal := curHeap[0]

		if parentIdx > 0 {
			parentVal = curHeap[parentIdx]
		}

		if curIdx*2+2 < len(curHeap) {
			lChildIdx := (curIdx * 2) + 1
			rChildIdx := (curIdx * 2) + 2

			lChildVal := curHeap[lChildIdx]
			rChildVal := curHeap[rChildIdx]

			if num > lChildVal || num > rChildVal {
				curHeap[parentIdx], curHeap[curIdx] = curHeap[curIdx], curHeap[parentIdx]
				curIdx = parentIdx
				continue
			}
		}

		if parentVal > num {
			curHeap[parentIdx], curHeap[curIdx] = curHeap[curIdx], curHeap[parentIdx]
			curIdx = parentIdx
			continue
		}
		break
	}

	*minBinHeap = curHeap
}
