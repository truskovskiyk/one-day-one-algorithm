package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func isOverlap(u Interval, v Interval) bool {
	if u.start < v.start {
		if u.end < v.start {
			return false
		}
	} else if v.start < u.start {
		if v.end < u.start {
			return false
		}
	}
	return true
}

func findOneLineTrackLength(oneLineTrack []Interval) int {
	sort.Slice(oneLineTrack, func(i, j int) bool { return oneLineTrack[i].start < oneLineTrack[j].start })

	stack := make([]Interval, 0)
	stack = append(stack, oneLineTrack[0])

	for i := 0; i < len(oneLineTrack); i++ {
		stackTop := stack[len(stack)-1]
		if isOverlap(stackTop, oneLineTrack[i]) {
			if stackTop.end < oneLineTrack[i].end {
				stackTop.end = oneLineTrack[i].end
				stack[len(stack)-1] = stackTop
			}
		} else {
			stack = append(stack, oneLineTrack[i])

		}
	}
	totalLength := 0
	for i := 0; i < len(stack); i++ {
		totalLength += stack[i].end - stack[i].start + 1
	}
	return totalLength
}

// Complete the gridlandMetro function below.
func GridlandMetro(n int, m int, k int, tracks [][]int) int {
	mapOfLineToTracks := make(map[int][]Interval)
	for i := 0; i < len(tracks); i++ {
		track := tracks[i]
		r := track[0]
		c1 := track[1]
		c2 := track[2]

		interval := Interval{c1, c2}
		mapOfLineToTracks[r] = append(mapOfLineToTracks[r], interval)
	}

	l := 0
	for _, oneLineTrack := range mapOfLineToTracks {
		l += findOneLineTrackLength(oneLineTrack)

	}
	r := n*m - l
	return r

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nmk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nmk[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	mTemp, err := strconv.ParseInt(nmk[1], 10, 64)
	checkError(err)
	m := int(mTemp)

	kTemp, err := strconv.ParseInt(nmk[2], 10, 64)
	checkError(err)
	k := int(kTemp)

	var track [][]int
	for i := 0; i < int(k); i++ {
		trackRowTemp := strings.Split(readLine(reader), " ")

		var trackRow []int
		for _, trackRowItem := range trackRowTemp {
			trackItemTemp, err := strconv.ParseInt(trackRowItem, 10, 64)
			checkError(err)
			trackItem := int(trackItemTemp)
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
