package main

import (
	"fmt"
	"sort"
	"bufio"
	"os"
	"io"
)

func BinarySearch(key int, arrayOfElements []int) int {
	start := 0
	end := len(arrayOfElements) - 1
	for start <= end {
		mid := (end + start) / 2
		if arrayOfElements[mid] == key {
			return mid
		}
		if arrayOfElements[mid] > key {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func findIndex(el int, c []int) int {
	for idx, element := range c {
		if element == el {
			return idx
		}
	}
	panic("can't find index")
}

func findIndex2(el int, idxOrigin int, c []int) int {
	for idx, element := range c {
		if element == el && idx != idxOrigin {
			return idx
		}
	}
	panic("can't find index")
}

func formatSolution(firstElement int, secondElement int, originArray []int) (int, int) {
	idxOrigin := findIndex(firstElement, originArray)
	kIdxOrigin := findIndex2(secondElement, idxOrigin, originArray)
	if idxOrigin < kIdxOrigin {
		return idxOrigin, kIdxOrigin
	} else {
		return kIdxOrigin, idxOrigin
	}

}
func IcecreamParlor(m int, c []int) (int, int) {
	originArray := make([]int, len(c))
	copy(originArray, c)
	sort.Ints(c)

	for idx, element := range c {
		k := m - element
		kIdx := BinarySearch(k, c)
		if kIdx != -1 && kIdx != idx {
			return formatSolution(element, k, originArray)
		}
	}
	panic("there is not solution")
}

func readFromStdIn() {
	readTask(bufio.NewReader(os.Stdin))
}

func readTask(ioReader io.Reader) {
	inSource := bufio.NewReader(ioReader)

	var numTests int
	fmt.Fscan(inSource, &numTests)

	for t := 0; t < numTests; t += 1 {
		var m, n, temp int
		fmt.Fscan(inSource, &m)
		fmt.Fscan(inSource, &n)
		c := make([]int, n)
		for i := 0; i < n; i += 1 {
			fmt.Fscan(inSource, &temp)
			c[i] = temp
		}
		i1, i2 := IcecreamParlor(m, c)
		fmt.Println(i1+1, i2+1)

	}
}

func main() {
	readFromStdIn()
}
