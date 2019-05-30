package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"math"
)


func minimumBribesResult(q []int) string {
	for i := len(q) - 1; i >= 0; i -= 1 {
		diff := q[i] - (1 + i)
		if diff > 2 {
			return "Too chaotic"
		}
	}


	moves := 0

	for i := len(q) - 1; i >= 0; i -= 1 {
		for j := int(math.Max(float64(q[i] - 2), 0)); j < i; j += 1 {
			if q[j] > q[i] {
				moves += 1
			}
		}

	}

	return strconv.Itoa(moves)
}

// Complete the minimumBribes function below.
func minimumBribes(q []int) {
	result := minimumBribesResult(q)
	fmt.Println(result)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
