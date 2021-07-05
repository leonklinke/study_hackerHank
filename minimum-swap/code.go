package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	swaps := 0
	searchingFrom := 0
	for {
		for i := searchingFrom; i < len(arr); i++ {

			if isInPosition(arr, i) {

				searchingFrom = i + 1

				if isLastElement(arr, i) {
					return int32(swaps)
				}
				continue
			}

			shifted := 0
			arr, shifted = swapThisPositions(arr, i)

			swaps += shifted
			break
		}
	}
}
func isInPosition(arr []int32, i int) bool {
	return arr[i] == int32(i+1)
}
func isLastElement(arr []int32, i int) bool {
	return len(arr) == i+1
}
func swapThisPositions(arr []int32, i int) ([]int32, int) {
	//[7,1,3,2,4...]
	// 0 <=> 3
	swaps := 1
	arr = swapPosition(arr, i, int(arr[i]-1))
	return arr, swaps
}

func swapPosition(arr []int32, actual, future int) []int32 {
	positionAux := arr[actual]
	arr[actual] = arr[future]
	arr[future] = positionAux
	return arr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
