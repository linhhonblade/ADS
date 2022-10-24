package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero.
	// Print the decimal value of each fraction on a new line with 6 places after the decimal.
	// Input format
	// The first line contains an integer, n, the size of the array.
	// The second line contains n space-separated integers that describe arr[n].
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	nTemp, _ := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	n := int32(nTemp)
	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}
	fmt.Println(arr)
	plusMinus(arr)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func plusMinus(arr []int32) {
	var posRatio, negRatio, zeroRatio float32
	for _, i := range arr {
		if i < 0 {
			negRatio++
		} else if i == 0 {
			zeroRatio++
		} else {
			posRatio++
		}
	}
	negRatio = negRatio / float32(len(arr))
	zeroRatio = zeroRatio / float32(len(arr))
	posRatio = posRatio / float32(len(arr))
	fmt.Printf("%.6f\n", posRatio)
	fmt.Printf("%.6f\n", negRatio)
	fmt.Printf("%.6f\n", zeroRatio)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
