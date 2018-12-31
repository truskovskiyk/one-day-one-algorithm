package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func intMod(num int, k int) int {
	mod := int(math.Mod(float64(num), float64(k)))
	return mod
}

func intMax(num int, k int) int {
	mod := int(math.Max(float64(num), float64(k)))
	return mod
}

func readTask(ioReader io.Reader) (int, int, []int) {
	inSource := bufio.NewReader(ioReader)

	var n, k, temp int
	fmt.Fscan(inSource, &n)
	fmt.Fscan(inSource, &k)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(inSource, &temp)
		s[i] = temp
	}
	return n, k, s
}
func readFromStdIn() (int, int, []int) {
	return readTask(bufio.NewReader(os.Stdin))
}

func readFromFile(filePath string) (int, int, []int) {
	f, _ := os.Open(filePath)
	defer f.Close()
	return readTask(bufio.NewReader(f))

}

func numberTheorySolution(k int, s []int) int {
	reminders := make([]int, k)
	for i := 0; i < len(s); i++ {
		r := intMod(s[i], k)
		reminders[r] += 1
	}
	result := 0
	topK := int(math.Ceil(float64(k) / 2))
	for i := 1; i < topK; i++ {
		if 2*i != k {
			result += intMax(reminders[i], reminders[k-i])
		}
	}
	// special case
	if intMod(k, 2) == 0 {
		if reminders[k/2] > 0 {
			result += 1
		}
	}
	// special case
	if reminders[0] > 0 {
		result += 1
	}
	return result
}
func main() {
	_, k, s := readFromStdIn()
	maxLen := numberTheorySolution(k, s)
	fmt.Println(maxLen)
}
