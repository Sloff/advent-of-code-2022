/*
Copyright © 2022 Ettienne Pitts

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package day3

import (
	"strings"
	"unicode"
)

func Day3Part1(data string) int {
	splitData := strings.Split(data, "\n")

	sum := 0

	for _, v := range splitData {
		value := strings.TrimSpace(v)
		size := int(len(value) / 2)

		for i := 0; i < size; i++ {
			if strings.ContainsRune(value[size:], rune(value[i])) {
				sum += charToValue(rune(value[i]))
				break
			}
		}
	}

	return sum
}

func Day3Part2(data string) int {
	splitData := strings.Split(data, "\n")

	sum := 0

	for i := 0; i < len(splitData); i += 3 {
		line1 := strings.TrimSpace(splitData[i])
		line2 := strings.TrimSpace(splitData[i+1])
		line3 := strings.TrimSpace(splitData[i+2])

		for i2 := 0; i2 < len(line1); i2++ {
			if strings.ContainsRune(line2, rune(line1[i2])) && strings.ContainsRune(line3, rune(line1[i2])) {
				sum += charToValue(rune(line1[i2]))
				break
			}
		}
	}

	return sum
}

func charToValue(char rune) int {
	if unicode.IsUpper(char) {
		return int(char) - 38
	} else {
		return int(char) - 96
	}
}
