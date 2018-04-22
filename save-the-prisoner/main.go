package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"math"
)

func intMod(num uint64, k uint64) uint64 {
	mod := uint64(math.Mod(float64(num), float64(k)))
	return mod
}

func readTask(ioReader io.Reader) {
	inSource := bufio.NewReader(ioReader)

	var t uint64
	fmt.Fscan(inSource, &t)
	for i := uint64(0); i < t; i++ {
		var numberOfPrisoners, numberOfSweets, chairNumberToStart uint64
		fmt.Fscan(inSource, &numberOfPrisoners)
		fmt.Fscan(inSource, &numberOfSweets)
		fmt.Fscan(inSource, &chairNumberToStart)
		r := savePrisoner(numberOfPrisoners, numberOfSweets, chairNumberToStart)
		fmt.Println(r)
	}
}

func readFromStdIn() {
	readTask(bufio.NewReader(os.Stdin))
}

func readFromFile(filePath string) {
	f, _ := os.Open(filePath)
	defer f.Close()
	readTask(bufio.NewReader(f))

}

func savePrisoner(numberOfPrisoners uint64, numberOfSweets uint64, chairNumberToStart uint64) uint64 {
	//intMod(chairNumberToStart, numberOfPrisoners)
	//intMod(numberOfSweets, numberOfPrisoners)
	pos := numberOfSweets + chairNumberToStart - 1
	ans := intMod(pos, numberOfPrisoners)
	if ans == 0 {
		ans = numberOfPrisoners
	}
	return ans

}
func main() {
	//filePath := "./DATA/in.txt"
	//readFromFile(filePath)
	readFromStdIn()
}
