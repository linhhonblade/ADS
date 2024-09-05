package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
	"unicode"
)

/*
https://adventofcode.com/2023/day/1
*/

func calibration_value(line string) int {
	var first, last rune
	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			first = runes[i]
			break
		}
	}
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			last = runes[i]
			break
		}
	}
	result, err := strconv.Atoi(string(first) + string(last))
	if err != nil {
		fmt.Printf("Failed to convert string to int: %s\n", err)
	}
	return result
}

func full_calibration_value(line string) int {
	runes := []rune(line)
	char3_map := map[string]int{
		"one": 1,
		"two": 2,
		"six": 6,
	}
	char4_map := map[string]int{
		"four": 4,
		"five": 5,
		"nine": 9,
	}
	char5_map := map[string]int{
		"three": 3,
		"seven": 7,
		"eight": 8,
	}
	var calibration_values []int
	i := 0
	for {
		if i >= len(runes) {
			break
		}
		if unicode.IsDigit(runes[i]) {
			calibration_values = append(calibration_values, int(runes[i])-48)
			i++
			continue
		}
		if i+3 <= len(runes) {
			if string(runes[i:i+3]) == "one" || string(runes[i:i+3]) == "two" || string(runes[i:i+3]) == "six" {
				calibration_values = append(calibration_values, char3_map[string(runes[i:i+3])])
				i += 2
				continue
			}
		}
		if i+4 <= len(runes) {
			if string(runes[i:i+4]) == "four" || string(runes[i:i+4]) == "five" || string(runes[i:i+4]) == "nine" {
				calibration_values = append(calibration_values, char4_map[string(runes[i:i+4])])
				i += 3
				continue
			}
		}
		if i+5 <= len(runes) {
			if string(runes[i:i+5]) == "three" || string(runes[i:i+5]) == "seven" || string(runes[i:i+5]) == "eight" {
				calibration_values = append(calibration_values, char5_map[string(runes[i:i+5])])
				i += 4
				continue
			}
		}
		i++
	}
	first := calibration_values[0]
	last := calibration_values[len(calibration_values)-1]
	firstStr := strconv.Itoa(first)
	lastStr := strconv.Itoa(last)
	concatStr := firstStr + lastStr
	result, err := strconv.Atoi(concatStr)
	if err != nil {
		fmt.Printf("Failed to convert string to int: %s\n", err)
	}
	return result
}

func file_calibration_value(calibration_file string) int {
	file, err := os.Open(calibration_file)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := 0
	for scanner.Scan() {
		result += calibration_value(scanner.Text())
	}
	return result
}

func file_full_calibration_value(calibration_file string) int {
	file, err := os.Open(calibration_file)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := 0
	for scanner.Scan() {
		result += full_calibration_value(scanner.Text())
	}
	return result
}

func TestCalibrationValue(t *testing.T) {
	t.Run("File sample.txt", func(t *testing.T) {
		calibration_file := "sample.txt"
		res := file_calibration_value(calibration_file)
		expected := 142
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("File calibration.txt", func(t *testing.T) {
		calibration_file := "calibration.txt"
		res := file_calibration_value(calibration_file)
		expected := 54605
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("File sample_2.txt - Full calibration", func(t *testing.T) {
		calibration_file := "sample_2.txt"
		res := file_full_calibration_value(calibration_file)
		expected := 281
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("File calibration.txt - Full calibration", func(t *testing.T) {
		calibration_file := "calibration.txt"
		res := file_full_calibration_value(calibration_file)
		expected := 55429
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
}
