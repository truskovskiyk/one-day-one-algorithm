package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the arrayManipulation function below.
func arrayManipulationBruteForce(n int32, queries [][]int32) int64 {
	array := make([]int, n)
	// O(n)
	for i := 0; i < len(array); i++ {
		array[i] = 0
	}

	// O(q)
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		start := query[0] - 1
		end := query[1] - 1
		value := query[2]
		// O(k)
		for j := start; j <= end; j ++ {
			array[j] = array[j] + int(value)
		}
	}
	fmt.Println(array)
	// O(n)
	maxValue := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > maxValue {
			maxValue = array[i]
		}
	}

	// O(n) + O(n) + O(q) * O(k) < O(n) + O(q * k) < O(n) + O(q * n)
	return int64(maxValue)
}


func arrayManipulationSolution(n int32, queries [][]int32) int64 {
	array := make([]int, n + 1)
	// O(n)
	for i := 0; i < len(array); i++ {
		array[i] = 0
	}

	// O(q)
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		start := query[0] - 1
		end := query[1] - 1
		value := query[2]

		array[start] += int(value)
		array[end + 1] -= int(value)
	}

	// O(n)
	maxValue := 0
	currentValue := 0
	for i := 0; i < len(array); i++ {
		currentValue += array[i]
		if currentValue > maxValue {
			maxValue = currentValue
		}
	}

	return int64(maxValue)
}

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {

	return arrayManipulationSolution(n, queries)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

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
