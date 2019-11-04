package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) int32 {
	all_possibles := [][]int32{
		{8, 1, 6, 3, 5, 7, 4, 9, 2},
		{6, 1, 8, 7, 5, 3, 2, 9, 4},
		{4, 9, 2, 3, 5, 7, 8, 1, 6},
		{2, 9, 4, 7, 5, 3, 6, 1, 8},
		{8, 3, 4, 1, 5, 9, 6, 7, 2},
		{4, 3, 8, 9, 5, 1, 2, 7, 6},
		{6, 7, 2, 1, 5, 9, 8, 3, 4},
		{2, 7, 6, 9, 5, 1, 4, 3, 8},
	}

	converted := convertTo1d(s)

	var min_cost int32 = math.MaxInt32

	var cost_loop int32

	for i := 0; i < 8; i++ {
		cost_loop = 0
		for j := 0; j < 9; j++ {
			cost_loop += absolute(converted[j] - all_possibles[i][j])
		}
		if cost_loop < min_cost {
			min_cost = cost_loop
		}
	}

	fmt.Println(min_cost)

	return 0
}

func convertTo1d(s [][]int32) []int32 {
	converted := make([]int32, 9)
	iter := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			converted[iter] = s[i][j]
			iter++
		}
	}
	return converted
}

func absolute(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, _ := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(readLine(reader), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

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
