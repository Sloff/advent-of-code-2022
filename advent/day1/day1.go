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
package day1

import (
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func setup(data string) []int {
	splitData := strings.Split(data, "\n")

	var index int = 1
	elves := map[int]int{}

	for _, v := range splitData {
		if _, ok := elves[index]; !ok {
			elves[index] = 0
		}

		if strings.TrimSpace(v) == "" {
			index++
		} else {
			val, err := strconv.Atoi(strings.TrimSpace(v))
			cobra.CheckErr(err)
			elves[index] += val
		}
	}

	var calories []int = []int{}

	for _, v := range elves {
		calories = append(calories, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	return calories
}

func Day1Part1(data string) int {
	return setup(data)[0]
}

func Day1Part2(data string) int {
	calories := setup(data)
	return calories[0] + calories[1] + calories[2]
}
