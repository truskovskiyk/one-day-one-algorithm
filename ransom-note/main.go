package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) string{
	magazineCount := make(map[string]int)

	// O(m)
	for i := 0; i < len(magazine); i++ {

		magazineValue := magazine[i]
		// O(1)
		if _, ok := magazineCount[magazineValue]; ok {
			magazineCount[magazineValue] += 1
			//do something here
		} else {
			// O(1) ~ O(m)
			magazineCount[magazineValue] = 1
		}
	}


	// O(n)
	for i := 0; i < len(note); i++ {

		noteValue := note[i]

		// O(1)
		if valueCount, ok := magazineCount[noteValue]; ok {
			if valueCount == 1 {
				delete(magazineCount, noteValue)
			} else {
				magazineCount[noteValue] -= 1
			}
		} else {
			return "No"
		}
	}


	return "Yes"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(readLine(reader), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(readLine(reader), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	result := checkMagazine(magazine, note)
	fmt.Println(result)
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
