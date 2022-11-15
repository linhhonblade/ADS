package main

import (
	"fmt"
)

/*
 * Complete the 'maximumPerimeterTriangle' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY sticks as parameter.
 */
func orderTriangle(triangle []int32) []int32 {
	max, min := triangle[0], triangle[0]
	var max_idx, min_idx int32
	for i, v := range triangle {
		if v > max {
			max = v
			max_idx = int32(i)
		}
		if v < min {
			min = v
			min_idx = int32(i)
		}
	}
	if min_idx == max_idx {
		return triangle
	} else {
		return []int32{triangle[min_idx], triangle[3-min_idx-max_idx], triangle[max_idx]}
	}
}
func maximumPerimeterTriangle(sticks []int32) []int32 {
	// Write your code here
	perimeter_map := make(map[int64][]int32)
	var perimeter, max_perimeter int64
	for i := 0; i < len(sticks); i++ {
		for j := i + 1; j < len(sticks); j++ {
			for k := j + 1; k < len(sticks); k++ {
				side1, side2, side3 := sticks[i], sticks[j], sticks[k]
				if side1+side2 > side3 && side1+side3 > side2 && side2+side3 > side1 {
					perimeter = int64(side1) + int64(side2) + int64(side3)
					perimeter_map[perimeter] = []int32{side1, side2, side3}
					if perimeter > max_perimeter {
						max_perimeter = perimeter
					}
				}
			}

		}
	}
	if max_perimeter == 0 {
		return []int32{-1}
	} else {
		return orderTriangle(perimeter_map[max_perimeter])
	}
}

func main() {
	arr := []int32{1000000000, 1000000000, 1000000000}
	// arr := []int32{1, 1, 1, 2, 3, 5}
	fmt.Println(maximumPerimeterTriangle(arr))
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	os.Setenv("OUTPUT_PATH", "test.txt")
// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	n := int32(nTemp)

// 	sticksTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

// 	var sticks []int32

// 	for i := 0; i < int(n); i++ {
// 		sticksItemTemp, err := strconv.ParseInt(sticksTemp[i], 10, 64)
// 		checkError(err)
// 		sticksItem := int32(sticksItemTemp)
// 		sticks = append(sticks, sticksItem)
// 	}

// 	result := maximumPerimeterTriangle(sticks)

// 	for i, resultItem := range result {
// 		fmt.Fprintf(writer, "%d", resultItem)

// 		if i != len(result)-1 {
// 			fmt.Fprintf(writer, " ")
// 		}
// 	}

// 	fmt.Fprintf(writer, "\n")

// 	writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
