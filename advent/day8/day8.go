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
package day8

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Tree struct {
	size int
	seen bool
}

func Day8Part1(data string) int {
	splitData := strings.Split(data, "\n")

	grid := [][]Tree{}

	for i, v := range splitData {
		grid = append(grid, []Tree{})

		for _, c := range strings.TrimSpace(v) {
			num, err := strconv.Atoi(string(c))
			cobra.CheckErr(err)
			grid[i] = append(grid[i], Tree{size: num, seen: false})
		}
	}

	result := 0

	// Seen from left & right
	for row := 0; row < len(grid); row++ {
		largestInRowFromLeft := grid[row][0].size
		largestInRowFromRight := grid[row][len(grid[row])-1].size

		for col := 0; col < len(grid[row]); col++ {
			inverseIndex := len(grid[row]) - 1 - col
			seenFromLeft := false
			seenFromRight := false
			if col == 0 {
				seenFromLeft = true
				seenFromRight = true
			}

			if grid[row][col].size > largestInRowFromLeft {
				seenFromLeft = true
				largestInRowFromLeft = grid[row][col].size
			}

			if grid[row][inverseIndex].size > largestInRowFromRight {
				seenFromRight = true
				largestInRowFromRight = grid[row][inverseIndex].size
			}

			if seenFromLeft {
				if !grid[row][col].seen {
					result++
				}

				grid[row][col].seen = true
			}

			if seenFromRight {
				if !grid[row][inverseIndex].seen {
					result++
				}

				grid[row][inverseIndex].seen = true
			}
		}
	}

	// Seen from top & bottom
	for col := 0; col < len(grid[0]); col++ {
		largestInColFromTop := grid[0][col].size
		largestInColFromBottom := grid[len(grid)-1][col].size

		for row := 0; row < len(grid); row++ {
			inverseIndex := len(grid) - 1 - row
			seenFromTop := false
			seenFromBottom := false
			if row == 0 {
				seenFromTop = true
				seenFromBottom = true
			}

			if grid[row][col].size > largestInColFromTop {
				seenFromTop = true
				largestInColFromTop = grid[row][col].size
			}

			if grid[inverseIndex][col].size > largestInColFromBottom {
				seenFromBottom = true
				largestInColFromBottom = grid[inverseIndex][col].size
			}

			if seenFromTop {
				if !grid[row][col].seen {
					result++
				}

				grid[row][col].seen = true
			}

			if seenFromBottom {
				if !grid[inverseIndex][col].seen {
					result++
				}

				grid[inverseIndex][col].seen = true
			}
		}
	}

	return result
}

func Day8Part2(data string) int {
	splitData := strings.Split(data, "\n")

	grid := [][]Tree{}

	for i, v := range splitData {
		grid = append(grid, []Tree{})

		for _, c := range strings.TrimSpace(v) {
			num, err := strconv.Atoi(string(c))
			cobra.CheckErr(err)
			grid[i] = append(grid[i], Tree{size: num, seen: false})
		}
	}

	highestScore := 0

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			amountToTop := 0

			for i := row - 1; i >= 0; i-- {
				amountToTop++

				if grid[i][col].size >= grid[row][col].size {
					break
				}
			}

			amountToBottom := 0

			for i := row + 1; i < len(grid); i++ {
				amountToBottom++

				if grid[i][col].size >= grid[row][col].size {
					break
				}
			}

			amountToLeft := 0

			for i := col - 1; i >= 0; i-- {
				amountToLeft++

				if grid[row][i].size >= grid[row][col].size {
					break
				}
			}

			amountToRight := 0

			for i := col + 1; i < len(grid[row]); i++ {
				amountToRight++

				if grid[row][i].size >= grid[row][col].size {
					break
				}
			}

			total := amountToBottom * amountToLeft * amountToTop * amountToRight

			if total > highestScore {
				highestScore = total
			}
		}
	}

	return highestScore
}
