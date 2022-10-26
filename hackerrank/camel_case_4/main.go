package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Input: "S;M;plasticCup()""
// Output ["S", "M", "plastic", "cup"]
func sanitizeInput(s string) string {
	var res []string
	var words []string
	var camel string
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "()")
	res = strings.Split(s, ";")
	if res[0] == "C" {
		res = append(res[:2], strings.Split(res[2], " ")...)
	} else {
		start := 0
		for i, c := range res[2] {
			if unicode.IsUpper(c) {
				if i != start {
					words = append(words, strings.ToLower(res[2][start:i]))
					start = i
				}
			}
		}
		words = append(words, strings.ToLower(res[2][start:]))
		res = append(res[:2], words...)
	}
	if res[0] == "S" {
		camel = strings.Join(res[2:], " ")
	} else {
		if res[1] == "C" {
			camel += leadUpper(res[2])
		} else {
			camel += res[2]
		}
		for _, w := range res[3:] {
			camel += leadUpper(w)
		}
		if res[1] == "M" {
			camel += "()"
		}
	}
	return camel
}

// Return the string with leading character uppercased
func leadUpper(s string) string {
	runeArr := []rune(s)
	runeArr[0] = unicode.ToUpper(runeArr[0])
	return string(runeArr)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		raw_input := scanner.Text()
		input := sanitizeInput(raw_input)
		fmt.Println(input)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
