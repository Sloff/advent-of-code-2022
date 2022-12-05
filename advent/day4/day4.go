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
package day4

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func Day4Part1(data string) int {
	splitData := strings.Split(data, "\n")

	result := 0

	for _, v := range splitData {
		pairs := strings.Split(strings.TrimSpace(v), ",")
		range1 := strings.Split(pairs[0], "-")
		range2 := strings.Split(pairs[1], "-")

		range1Lower, err := strconv.Atoi(range1[0])
		cobra.CheckErr(err)
		range1Higher, err := strconv.Atoi(range1[1])
		cobra.CheckErr(err)

		range2Lower, err := strconv.Atoi(range2[0])
		cobra.CheckErr(err)
		range2Higher, err := strconv.Atoi(range2[1])
		cobra.CheckErr(err)

		if (range1Lower >= range2Lower && range1Higher <= range2Higher) || (range2Lower >= range1Lower && range2Higher <= range1Higher) {
			result++
		}
	}

	return result
}

func Day4Part2(data string) int {
	splitData := strings.Split(data, "\n")

	result := 0

	for _, v := range splitData {
		pairs := strings.Split(strings.TrimSpace(v), ",")
		range1 := strings.Split(pairs[0], "-")
		range2 := strings.Split(pairs[1], "-")

		range1Lower, err := strconv.Atoi(range1[0])
		cobra.CheckErr(err)
		range1Higher, err := strconv.Atoi(range1[1])
		cobra.CheckErr(err)

		range2Lower, err := strconv.Atoi(range2[0])
		cobra.CheckErr(err)
		range2Higher, err := strconv.Atoi(range2[1])
		cobra.CheckErr(err)

		if (range1Lower >= range2Lower && range1Lower <= range2Higher) || (range1Higher >= range2Lower && range1Higher <= range2Higher) || (range2Lower >= range1Lower && range2Lower <= range1Higher) || (range2Higher >= range1Lower && range2Higher <= range1Higher) {
			result++
		}
	}

	return result
}
