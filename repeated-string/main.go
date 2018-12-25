package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func letterInString(a string, s string) int{
	count := 0
	for _, char := range s {
		if a == string(char) {
			count += 1
		}
	}
	return count
}

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	stringLen := int64(len(s))
	numRepeat := n / stringLen
	fmt.Println(numRepeat)

	disLen := n - numRepeat * stringLen

	letterToFind := "a"
	countInFullString := int64(letterInString(letterToFind, s)) * numRepeat
	countInRest := int64(letterInString(letterToFind, s[:disLen]))

	return countInFullString + countInRest
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

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
