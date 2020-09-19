package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var applesLandInRange int32 = 0
	var orangesLandInRange int32 = 0

	for _, apple := range apples {
		if landInRegion(s, t, a+apple) {
			applesLandInRange++
		}
	}

	for _, orange := range oranges {
		if landInRegion(s, t, b+orange) {
			orangesLandInRange++
		}
	}
	fmt.Println(applesLandInRange)
	fmt.Println(orangesLandInRange)
}

func landInRegion(s int32, t int32, fruitPosition int32) bool {
	if fruitPosition >= s && fruitPosition <= t {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	st := strings.Split(readLine(reader), " ")

	sTemp, _ := strconv.ParseInt(st[0], 10, 64)
	// checkError(err)
	s := int32(sTemp)

	tTemp, _ := strconv.ParseInt(st[1], 10, 64)
	// checkError(err)
	t := int32(tTemp)

	ab := strings.Split(readLine(reader), " ")

	aTemp, _ := strconv.ParseInt(ab[0], 10, 64)
	// checkError(err)
	a := int32(aTemp)

	bTemp, _ := strconv.ParseInt(ab[1], 10, 64)
	// checkError(err)
	b := int32(bTemp)

	mn := strings.Split(readLine(reader), " ")

	mTemp, _ := strconv.ParseInt(mn[0], 10, 64)
	// checkError(err)
	m := int32(mTemp)

	nTemp, _ := strconv.ParseInt(mn[1], 10, 64)
	// checkError(err)
	n := int32(nTemp)

	applesTemp := strings.Split(readLine(reader), " ")

	var apples []int32

	for i := 0; i < int(m); i++ {
		applesItemTemp, _ := strconv.ParseInt(applesTemp[i], 10, 64)
		// checkError(err)
		applesItem := int32(applesItemTemp)
		apples = append(apples, applesItem)
	}

	orangesTemp := strings.Split(readLine(reader), " ")

	var oranges []int32

	for i := 0; i < int(n); i++ {
		orangesItemTemp, _ := strconv.ParseInt(orangesTemp[i], 10, 64)
		// checkError(err)
		orangesItem := int32(orangesItemTemp)
		oranges = append(oranges, orangesItem)
	}

	countApplesAndOranges(s, t, a, b, apples, oranges)
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
