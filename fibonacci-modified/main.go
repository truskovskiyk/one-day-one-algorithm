package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"math/big"
)


func readFromStdIn() (int64, int64, int64) {
	return readTask(bufio.NewReader(os.Stdin))
}

func readFromFile(filePath string) (int64, int64, int64) {
	f, _ := os.Open(filePath)
	defer f.Close()
	return readTask(bufio.NewReader(f))
}

func readTask(ioReader io.Reader) (int64, int64, int64) {
  inSource := bufio.NewReader(ioReader)

  var n, t1, t2 int64
	fmt.Fscan(inSource, &t1)
	fmt.Fscan(inSource, &t2)
	fmt.Fscan(inSource, &n)

  return n, t1, t2
}

func fibonacciModified(n int64, t1 big.Int, t2 big.Int) big.Int {
	result := t2
	for i := int64(2); i < n; i += 1 {
		fmt.Println(result)
		// b.Mul(b, b)
		// result.Add(t1, b)
		// fmt.Println(result)
		// t1 = t2
		// t2 = result
	}
	return result
}
func main() {
	// n, t1, t2 := readFromStdIn()

	filePath := "./input.txt"
	n, t1, t2 := readFromFile(filePath)
	fmt.Println(n, t1, t2)
	result := fibonacciModified(n, *big.NewInt(t1), *big.NewInt(t2))
	fmt.Println(string(*result))
}
