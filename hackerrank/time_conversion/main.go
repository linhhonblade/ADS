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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	time_raw := s[:len(s)-2]
	notation := s[len(s)-2:]
	hour, _ := strconv.ParseInt((s[:2]), 10, 8)
	if notation == "AM" && hour == 12 {
		time_raw = "00" + time_raw[2:]
	} else if notation == "PM" && hour != 12 {
		hour += 12
		time_raw = fmt.Sprintf("%d", hour) + time_raw[2:]
	}
	return time_raw
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	os.Setenv("OUTPUT_PATH", "test.txt")

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
