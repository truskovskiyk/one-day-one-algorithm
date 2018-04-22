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
	year  int
}

type ArrayOfItems []Item

func (a ArrayOfItems) Len() int      { return len(a) }
func (a ArrayOfItems) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ArrayOfItems) Less(i, j int) bool {
	return a[i].price < a[j].price
}

func readFromStdIn() (int, []Item) {
	return readTask(bufio.NewReader(os.Stdin))
}

func readFromFile(filePath string) (int, []Item) {
	f, _ := os.Open(filePath)
	defer f.Close()
	return readTask(bufio.NewReader(f))
}

func readTask(ioReader io.Reader) (int, []Item) {
	inSource := bufio.NewReader(ioReader)

	var n, temp int
	fmt.Fscan(inSource, &n)

	p := make([]Item, n)
	for i := int(0); i < n; i++ {
		fmt.Fscan(inSource, &temp)
		p[i] = Item{price: temp, year: i}

	}
	return n, p
}

func solveTaskSecond(numberDays int, prices []Item) int {
	sort.Sort(ArrayOfItems(prices))
	minLoss := math.Inf(+1)
	for i := int(0); i < numberDays-1; i++ {
		if prices[i].year > prices[i+1].year {
			//loss := prices[i].price - prices[i+1].price
			loss := prices[i+1].price - prices[i].price
			if float64(loss) < minLoss {
				minLoss = float64(loss)
			}
		}
	}
	return int(minLoss)
}

func main() {
	//fileName := "./DATA/in.txt"
	//n, p := readFromFile(fileName)

	n, p := readFromStdIn()
	minLoss := solveTaskSecond(n, p)
	fmt.Println(minLoss)

}
