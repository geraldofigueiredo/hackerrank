package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var maxValue int64

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	array := make([]int64, n)

	for _, query := range queries {
		applyQuery(query, &array)
	}

	return maxValue
}

func applyQuery(query []int32, array *[]int64) {
	start, end, acc := queryConverter(query)
	for start < end {
		(*array)[start] += acc
		if (*array)[start] > maxValue {
			maxValue = (*array)[start]
		}
		start++
	}
}

func queryConverter(query []int32) (int32, int32, int64) {
	start := query[0] - 1
	end := query[1]
	var acc int64 = int64(query[2])

	return start, end, acc
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, _ := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, _ := strconv.ParseInt(nm[0], 10, 64)
	// checkError(err)
	n := int32(nTemp)

	mTemp, _ := strconv.ParseInt(nm[1], 10, 64)
	// checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, _ := strconv.ParseInt(queriesRowItem, 10, 64)
			// checkError(err)
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
