package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"math"
	"sort"
)

type Item struct {
	price int
	year int
}
func readFromStdIn() (int, []int) {
	return readTask(bufio.NewReader(os.Stdin))
}


func readFromFile(filePath string) (int, []int) {
	f, _ := os.Open(filePath)
	defer f.Close()
	return readTask(bufio.NewReader(f))
}


func solveTaskFirst(numberDays int, prices []int) int{
	//fmt.Println("numberDays", numberDays)
	//fmt.Println("prices", prices)
	minLoss := math.Inf(+1)
	for i := int(0); i < numberDays; i++ {
		for j := i + 1; j < numberDays; j++ {
			b := prices[i]
			c := prices[j]
			loss := b - c
			if c >= b {
				continue
			}
			//fmt.Println("duy", i, "sell", j, "loss", loss)
			//fmt.Println(c)
			if float64(loss) < minLoss {
				minLoss = float64(loss)
			}
		}
	}
	return int(minLoss)
}

func readTask(ioReader io.Reader) (int, []int) {
	inSource := bufio.NewReader(ioReader)

	var n int
	fmt.Fscan(inSource, &n)

	p := make([]int, n)
	for i := int(0); i < n; i++ {
		fmt.Fscan(inSource, &p[i])
	}
	return n, p
}

func solveTaskSecond(numberDays int, prices []int) int{
	items :=
	sort.Ints(prices)
	fmt.Println(prices)
	//fmt.Println("numberDays", numberDays)
	//fmt.Println("prices", prices)
	minLoss := math.Inf(+1)
	for i := int(0); i < numberDays; i++ {
		for j := i + 1; j < numberDays; j++ {
			b := prices[i]
			c := prices[j]
			loss := b - c
			if c >= b {
				continue
			}
			//fmt.Println("duy", i, "sell", j, "loss", loss)
			//fmt.Println(c)
			if float64(loss) < minLoss {
				minLoss = float64(loss)
			}
		}
	}
	return int(minLoss)
}

func main() {
	fileName := "./DATA/in.txt"
	n, p := readFromFile(fileName)



	//n, p := readFromStdIn()
	minLoss := solveTaskSecond(n, p)
	fmt.Println(minLoss)


}
