package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
	"unicode"
)

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
}
