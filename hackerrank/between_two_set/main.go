package main

import (
	"fmt"
)

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var count int32
	var valid bool
	max, min := a[0], b[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	for _, v := range b {
		if v < min {
			min = v
		}
	}
	for i := max; i <= min; i++ {
		valid = true
		for _, v := range a {
			if i%v != 0 {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		for _, v := range b {
			if v%i != 0 {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
	}
	return count
}

func main() {
	a := []int32{2, 4}
	b := []int32{16, 32, 96}
	res := getTotalX(a, b)
	fmt.Println(res)
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//     checkError(err)

//     defer stdout.Close()

//     writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

//     firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

//     nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
//     checkError(err)
//     n := int32(nTemp)

//     mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
//     checkError(err)
//     m := int32(mTemp)

//     arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

//     var arr []int32

//     for i := 0; i < int(n); i++ {
//         arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
//         checkError(err)
//         arrItem := int32(arrItemTemp)
//         arr = append(arr, arrItem)
//     }

//     brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

//     var brr []int32

//     for i := 0; i < int(m); i++ {
//         brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
//         checkError(err)
//         brrItem := int32(brrItemTemp)
//         brr = append(brr, brrItem)
//     }

//     total := getTotalX(arr, brr)

//     fmt.Fprintf(writer, "%d\n", total)

//     writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }
