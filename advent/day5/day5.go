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
package day5

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func getInstructionsAndStacks(data string) ([][]string, map[int]string) {
	splitData := strings.Split(data, "\n")

	stacks := []string{}
	instructionStrings := []string{}
	newlineFound := false

	for _, v := range splitData {
		// value := strings.TrimSpace(v)

		if !newlineFound {
			if strings.TrimSpace(v) == "" {
				newlineFound = true
			} else {
				stacks = append(stacks, v)
			}
		} else {
			instructionStrings = append(instructionStrings, v)
		}
	}

	stackMap := map[int]string{}

	for _, v := range stacks[:len(stacks)-1] {
		col := 1
		for i := 1; i < len(v)-1; i += 4 {
			if v[i] != byte(' ') {
				if _, ok := stackMap[col]; !ok {
					stackMap[col] = ""
				}

				stackMap[col] = stackMap[col] + string(v[i])
			}
			col++
		}
	}

	instructions := [][]string{}
	r := regexp.MustCompile(`move (.*) from (.*) to (.*)$`)

	for _, v := range instructionStrings {
		instructions = append(instructions, r.FindStringSubmatch(v)[1:])
	}

	return instructions, stackMap
}

func Day5Part1(data string) string {
	instructions, stackMap := getInstructionsAndStacks(data)

	for _, instruction := range instructions {
		iter, err := strconv.Atoi(strings.TrimSpace(instruction[0]))
		cobra.CheckErr(err)
		from, err := strconv.Atoi(strings.TrimSpace(instruction[1]))
		cobra.CheckErr(err)
		to, err := strconv.Atoi(strings.TrimSpace(instruction[2]))
		cobra.CheckErr(err)

		for i := 0; i < iter; i++ {
			valueToMove := stackMap[from][0]
			stackMap[from] = stackMap[from][1:]
			stackMap[to] = string(valueToMove) + stackMap[to]
		}
	}

	result := ""

	for i := 1; i <= len(stackMap); i++ {
		result = result + string(stackMap[i][0])
	}

	return result
}

func Day5Part2(data string) string {
	instructions, stackMap := getInstructionsAndStacks(data)

	for _, instruction := range instructions {
		amount, err := strconv.Atoi(strings.TrimSpace(instruction[0]))
		cobra.CheckErr(err)
		from, err := strconv.Atoi(strings.TrimSpace(instruction[1]))
		cobra.CheckErr(err)
		to, err := strconv.Atoi(strings.TrimSpace(instruction[2]))
		cobra.CheckErr(err)

		valueToMove := stackMap[from][0:amount]
		stackMap[from] = stackMap[from][amount:]
		stackMap[to] = valueToMove + stackMap[to]
	}

	result := ""

	for i := 1; i <= len(stackMap); i++ {
		result = result + string(stackMap[i][0])
	}

	return result
}
