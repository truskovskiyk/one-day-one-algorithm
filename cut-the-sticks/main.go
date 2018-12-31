package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func min(arr []int32) int32 {
	minElement := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < minElement {
			minElement = arr[i]
		}
	}
	return minElement
}

func insertIntoStart(arr []int32, element int32) []int32 {
	arr = append(arr, 0)
	copy(arr[0+1:], arr[0:])
	arr[0] = element
	return arr
}

// Complete the cutTheSticks function below.
func CutTheSticks(arr []int32) []int32 {
	if len(arr) == 0 {
		return make([]int32, 0)
	}

	minElement := min(arr)
	newArray := make([]int32, 0)
	for i := 0; i < len(arr); i++ {
		el := arr[i] - minElement

		if el > 0 {
			newArray = append(newArray, el)

		}
	}
	res := CutTheSticks(newArray)
	res = insertIntoStart(res, int32(len(arr)))
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := CutTheSticks(arr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
