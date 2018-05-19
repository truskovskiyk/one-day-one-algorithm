package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


type Graph struct {
	numVertex int32
	numEdges  int32
	adj [][]int32
}

func (g Graph) addEdge(v int32, w int32) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
}


func shortedPathFromedgeTo(edgeTo []int32, target int32, source int32) []int32 {
	path := make([]int32, 0)
	for x := target; x != source; x = edgeTo[x] {
		path = append(path, x)
	}
	return path
}
func (g Graph) BFS (source int32) [][]int32 {
	marked := make([]bool, g.numVertex)
	queue := make([]int32, 0)
	edgeTo := make([]int32, g.numVertex)
	marked[source] = true
	queue = append(queue, source)
	for ; len(queue) > 0; {
		v := queue[0]
		queue = queue[1:]
		for _, w := range g.adj[v] {
			if marked[w] == false {
				marked[w] = true
				edgeTo[w] = v
				queue = append(queue, w)
			}
		}

	}
	// find shorted path
	shortedPath := make([][]int32, g.numVertex)
	for targer := int32(0); targer < g.numVertex; targer += 1 {
		shortedPath[targer] = make([]int32, 0)
		if marked[targer] {
			shortedPath[targer] = shortedPathFromedgeTo(edgeTo, targer, source)
		}
	}

	return shortedPath
}

func NewGraph(n int32, m int32, edges [][]int32) Graph {
	graph := new(Graph)
	graph.numVertex = n
	graph.numEdges = m
	graph.adj = make([][]int32, n)
	for i := int32(0); i < n; i += 1 {
		graph.adj[i] = make([]int32, 0)
	}
	for _, e := range edges {
		graph.addEdge(e[0], e[1])
	}
	return *graph
}

func formatResult(numVertex int32, shortedPath [][]int32, source int32) []int32 {
	results := make([]int32, 0)
	for i := int32(0); i < numVertex; i += 1 {
		if i ==  source  {
			continue
		}
		dist := int32(len(shortedPath[i]) * 6)
		if dist == 0 {
			dist = -1
		}
		results = append(results, dist)
	}
	return results
}

// Complete the bfs function below.
func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	// start index from 0
	for _, e := range edges {
		e[0] = e[0] - 1
		e[1] = e[1] - 1
	}
	source := s - 1
	graph := NewGraph(n,m,edges)
	shortedPath := graph.BFS(source)
	return formatResult(graph.numVertex, shortedPath, source)

}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	filePath := "bfs_input.txt"
	f, _ := os.Open(filePath)
	defer f.Close()
	reader := bufio.NewReader(f)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nm := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nm[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nm[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		var edges [][]int32
		for i := 0; i < int(m); i++ {
			edgesRowTemp := strings.Split(readLine(reader), " ")

			var edgesRow []int32
			for _, edgesRowItem := range edgesRowTemp {
				edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
				checkError(err)
				edgesItem := int32(edgesItemTemp)
				edgesRow = append(edgesRow, edgesItem)
			}

			if len(edgesRow) != int(2) {
				panic("Bad input")
			}

			edges = append(edges, edgesRow)
		}

		sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		s := int32(sTemp)

		result := bfs(n, m, edges, s)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result) - 1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

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
