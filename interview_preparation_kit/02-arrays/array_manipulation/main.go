package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type query struct {
	start int32
	end   int32
	value int64
}

func (q *query) update(start, end int32, value int64, sumFlag bool) {
	q.start = start
	q.end = end

	if sumFlag {
		q.value += value
	} else {
		q.value = value
	}
}

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	var maxValue query

	array := make([]int64, n)

	// processing the first query
	start, end, value := queryConverter(queries[0])
	maxValue.update(start, end, value, false)
	queries = queries[1:]
	applyQuery(start, end, value, &array)
	fmt.Println("\n", array)

	for i := range queries {
		start, end, value := queryConverter(queries[i])
		applyQuery(start, end, value, &array)
		fmt.Println("\n", array)

		// out of maxValue range
		if end < maxValue.start || start > maxValue.end {
			fmt.Print("1", maxValue, start, end, value)
			if value > maxValue.value {
				maxValue.update(start, end, value, false)
			}
			fmt.Println(" -> ", maxValue)
			continue
		}

		// in range
		if start >= maxValue.start && end <= maxValue.end {
			fmt.Print("2", maxValue, start, end, value)
			maxValue.update(start, end, value, true)
			fmt.Println(" -> ", maxValue)
			continue
		}

		// out of range
		if start < maxValue.start && end <= maxValue.end {
			fmt.Print("3", maxValue, start, end, value)
			maxValue.update(maxValue.start, end, value, true)
			fmt.Println(" -> ", maxValue)
			continue
		}

		if start >= maxValue.start && end > maxValue.end {
			fmt.Print("4", maxValue, start, end, value)
			maxValue.update(start, maxValue.end, value, true)
			fmt.Println(" -> ", maxValue)
			continue
		}

		if start <= maxValue.start && end >= maxValue.end {
			fmt.Print("5", maxValue, start, end, value)
			maxValue.update(maxValue.start, maxValue.end, value, true)
			fmt.Println(" -> ", maxValue)
			continue
		}
	}

	fmt.Println("\n\n", maxValue, max)

	fmt.Println(array)

	return maxValue.value
}

func queryConverter(query []int32) (int32, int32, int64) {
	start := query[0]
	end := query[1]
	var value int64 = int64(query[2])

	return start, end, value
}

var max int64

func applyQuery(start, end int32, value int64, array *[]int64) {
	for i := start; i < end; i++ {
		(*array)[i] += value
		if (*array)[i] > max {
			max = (*array)[i]
		}
	}
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
