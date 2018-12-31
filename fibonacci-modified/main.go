package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
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

func FibonacciModified(n int64, t1 int64, t2 int64) string {
	t11 := big.NewInt(t1)
	t22 := big.NewInt(t2)
	result := big.NewInt(t2)
	for i := int64(2); i < n; i += 1 {
		b := big.NewInt(0)
		b.Set(t22)
		b = b.Mul(b, b)
		result.Add(t11, b)
		t11.Set(t22)
		t22.Set(result)
	}
	return result.String()
}
func main() {
	n, t1, t2 := readFromStdIn()

	//filePath := "./input.txt"
	//n, t1, t2 := readFromFile(filePath)
	result := FibonacciModified(n, t1, t2)
	fmt.Println(result)
}
