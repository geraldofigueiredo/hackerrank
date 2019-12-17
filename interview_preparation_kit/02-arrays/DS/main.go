package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var max int32 = -math.MaxInt32
	var temp int32 = -math.MaxInt32

	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[i])-2; j++ {
			temp = arr[i][j] + arr[i][j+1] + arr[i][j+2]
			temp += arr[i+1][j+1]
			temp += arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			if temp > max {
				max = temp
			}
		}
	}

	return max
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, _ := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	// for i := 0; i < 6; i++ {
	// 	arrRowTemp := strings.Split(readLine(reader), " ")

	// 	var arrRow []int32
	// 	for _, arrRowItem := range arrRowTemp {
	// 		arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
	// 		checkError(err)
	// 		arrItem := int32(arrItemTemp)
	// 		arrRow = append(arrRow, arrItem)
	// 	}

	// 	if len(arrRow) != int(6) {
	// 		panic("Bad input")
	// 	}

	// 	arr = append(arr, arrRow)
	// }

	arr = make([][]int32, 6)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int32, 6)
	}

	// arr = [][]int32{
	// 	{1, 2, 3, 4, 5, 6},
	// 	{6, 5, 4, 3, 2, 1},
	// 	{5, 6, 7, 8, 9, 0},
	// 	{0, 9, 8, 7, 6, 5},
	// 	{1, 3, 5, 7, 9, 0},
	// 	{2, 4, 6, 8, 0, 9},
	// }

	arr = [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
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
