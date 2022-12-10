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
package day10

import (
	"math"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func checkCycle(i int) bool {
	return i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220
}

func Day10Part1(data string) int {
	splitData := strings.Split(data, "\n")

	cycle := 1
	registerX := 1

	result := 0

	for _, v := range splitData {
		val := strings.TrimSpace(v)

		if checkCycle(cycle) {
			result += cycle * registerX
		}

		if val == "noop" {
			cycle++
		} else {
			instruction := strings.Split(val, " ")
			amount, err := strconv.Atoi(instruction[1])
			cobra.CheckErr(err)

			cycle++
			if checkCycle(cycle) {
				result += cycle * registerX
			}
			cycle++
			registerX += amount
		}
	}

	return result
}

func checkCycle2(i int) bool {
	return i%40 == 0
}

func Day10Part2(data string) string {
	splitData := strings.Split(data, "\n")

	cycle := 0
	registerX := 1
	output := ""

	for _, v := range splitData {
		val := strings.TrimSpace(v)

		if checkCycle2(cycle) {
			output = output + "\n"
		}

		crtDrawPoint := cycle % 40

		if math.Abs(float64(crtDrawPoint)-float64(registerX)) <= 1 {
			output = output + "#"
		} else {
			output = output + "."
		}

		if val == "noop" {
			cycle++
		} else {
			instruction := strings.Split(val, " ")
			amount, err := strconv.Atoi(instruction[1])
			cobra.CheckErr(err)

			cycle++
			if checkCycle2(cycle) {
				output = output + "\n"
			}
			crtDrawPoint = cycle % 40

			if math.Abs(float64(crtDrawPoint)-float64(registerX)) <= 1 {
				output = output + "#"
			} else {
				output = output + "."
			}
			cycle++
			registerX += amount
		}
	}

	return output
}
