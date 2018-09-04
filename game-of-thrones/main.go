package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func even(number int) bool {
	return number%2 == 0
}

func odd(number int) bool {
	return !even(number)
}

func getCounter(s string) map[rune]int {
	counter := make(map[rune]int)
	for _, c := range s {
		if _, ok := counter[c]; ok {
			counter[c] += 1
		} else {
			counter[c] = 1
		}
	}
	return counter
}
func gameOfThrones(s string) string {
	counter := getCounter(s)
	numAllowed := 1
	for _, v := range counter {
		if odd(v) {
			numAllowed -= 1
		}
	}
	if numAllowed < 0 {
		return "NO"
	}
	return "YES"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := gameOfThrones(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
