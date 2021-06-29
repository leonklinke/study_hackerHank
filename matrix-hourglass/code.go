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

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

var matrixOrder = int32(6)

func hourglassSum(arr [][]int32) int32 {
	// going from top to down
	highestSum := int32(math.Inf(-1))
	for i := 0; i < int(matrixOrder-2); i++ {
		for j := 0; j < int(matrixOrder-2); j++ {
			sum := getHourGlassSum(arr, i, j)
			highestSum = verifyIfIsBigger(sum, highestSum)
		}
	}
	// navigateHorizontal(arr)
	return highestSum
}
func getHourGlassSum(arr [][]int32, i, j int) int32 {
	a := arr[i][j]
	b := arr[i][j+1]
	c := arr[i][j+2]
	d := arr[i+1][j+1]
	e := arr[i+2][j]
	f := arr[i+2][j+1]
	g := arr[i+2][j+2]
	return a + b + c + d + e + f + g
}

func verifyIfIsBigger(sum, highestSum int32) int32 {
	if sum > highestSum {
		return sum
	}
	return highestSum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
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
