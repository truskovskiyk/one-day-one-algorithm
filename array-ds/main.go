package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func getHourglass(arr [][]int, i int, j int) int {
	a := arr[i][j]
	b := arr[i][j+1]
	c := arr[i][j+2]

	d := arr[i+1][j+1]

	e := arr[i+2][j]
	f := arr[i+2][j+1]
	g := arr[i+2][j+2]

	return a + b + c + d + e + f + g
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int) int {
	maxHourglassSum := MinInt
	for i := 0; i < len(arr)-2; i += 1 {
		for j := 0; j < len(arr)-2; j += 1 {
			currentHourglassSum := getHourglass(arr, i, j)
			if currentHourglassSum > maxHourglassSum {
				maxHourglassSum = currentHourglassSum
			}
		}
	}

	return maxHourglassSum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

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
