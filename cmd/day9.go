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
package cmd

import (
	"fmt"

	"github.com/Sloff/advent-of-code-2022/advent/day9"
	"github.com/Sloff/advent-of-code-2022/utils"
	"github.com/spf13/cobra"
)

// day9Cmd represents the day9 command
var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "Day 9 of advent of code",
	Run: func(cmd *cobra.Command, args []string) {
		data := utils.GetData("day9")
		fmt.Println(day9.Day9Part1(data))
		fmt.Println(day9.Day9Part2(data))
	},
}

func init() {
	rootCmd.AddCommand(day9Cmd)
}
