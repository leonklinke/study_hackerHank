package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

const MAXBRIBE = 2

func minimumBribes(q []int32) {
	// Write your code here
	totalShifts := 0
	higherSticker := int32(len(q))
	for {
		if higherSticker == 1 {
			fmt.Println(totalShifts)
			return
		}

		if isInRightPlace(q, higherSticker) {
			higherSticker--
			continue
		}
		position := getHigherStickerPossiblePosition(q, higherSticker)
		if position == -1 {
			fmt.Println("Too chaotic")
			return
		}
		higherSticker--

		bribes := 0
		q, bribes = transferToTheRightPlace(q, position)

		totalShifts += bribes
	}
}
func isInRightPlace(q []int32, stickerNumber int32) bool {
	return q[stickerNumber-1] == stickerNumber
}
func getHigherStickerPossiblePosition(q []int32, higherStick int32) int {
	for i := 0; i <= MAXBRIBE; i++ {
		if q[int(higherStick)-1-i] == higherStick {
			return int(higherStick) - 1 - i
		}
	}
	return -1
}

func transferToTheRightPlace(q []int32, positionInArray int) ([]int32, int) {

	bribesNeeded := getBribesNeeded(q, positionInArray)

	q = shiftRight(q, positionInArray, bribesNeeded)

	return q, int(bribesNeeded)
}

func getBribesNeeded(q []int32, positionInArray int) int {
	stickerNumber := q[positionInArray]
	liesAt := positionInArray + 1
	return int(stickerNumber) - liesAt
}

func shiftRight(q []int32, positionInArray, shifts int) []int32 {
	actualSticker := q[positionInArray]
	q[positionInArray] = q[positionInArray+1]
	if shifts > 1 {
		q[positionInArray+1] = q[positionInArray+2]
		q[positionInArray+2] = actualSticker
		return q
	}
	q[positionInArray+1] = actualSticker
	return q
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
