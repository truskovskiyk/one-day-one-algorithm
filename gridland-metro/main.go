package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sort"
)


// Complete the gridlandMetro function below.
func GridlandMetro(n int32, m int32, k int32, tracks [][]int32) int32 {

	sort.Slice(tracks, func(i, j int) bool { return tracks[i][0] < tracks[j][0] })

	counter := make(map[int32][][]int32)
	for i := 0; i < len(tracks); i++ {

		track := tracks[i]
		fmt.Println(track)
		r := track[0]
		c1 := track[1]
		c2 := track[2]
		counter[r] = append(counter[r], []int32{c1, c2})
	}
	fmt.Println(counter)
	return int32(1)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nmk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nmk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nmk[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	kTemp, err := strconv.ParseInt(nmk[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	var track [][]int32
	for i := 0; i < int(k); i++ {
		trackRowTemp := strings.Split(readLine(reader), " ")

		var trackRow []int32
		for _, trackRowItem := range trackRowTemp {
			trackItemTemp, err := strconv.ParseInt(trackRowItem, 10, 64)
			checkError(err)
			trackItem := int32(trackItemTemp)
			trackRow = append(trackRow, trackItem)
		}

		if len(trackRow) != int(3) {
			panic("Bad input")
		}

		track = append(track, trackRow)
	}

	result := GridlandMetro(n, m, k, track)

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
