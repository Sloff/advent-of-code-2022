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
package day2

import (
	"strings"
)

func Day2Part1(data string) int {
	splitData := strings.Split(data, "\n")
	score := 0

	for _, v := range splitData {
		values := strings.Split(v, " ")

		if len(values) < 2 {
			break
		}

		values[0] = strings.TrimSpace(values[0])
		values[1] = strings.TrimSpace(values[1])

		if values[0] == "A" { // Rock
			if values[1] == "Y" { // Paper
				score += 2 + 6
			} else if values[1] == "Z" { // Scissors
				score += 3
			} else { // Rock
				score += 1 + 3
			}
		} else if values[0] == "B" { // Paper
			if values[1] == "Z" { // Scissors
				score += 3 + 6
			} else if values[1] == "X" { // Rock
				score += 1
			} else { // Paper
				score += 2 + 3
			}
		} else if values[0] == "C" { // Scissors
			if values[1] == "X" { // Rock
				score += 1 + 6
			} else if values[1] == "Y" { // Paper
				score += 2
			} else { // Scissors
				score += 3 + 3
			}
		}
	}

	return score
}

func Day2Part2(data string) int {
	splitData := strings.Split(data, "\n")
	score := 0

	for _, v := range splitData {
		values := strings.Split(v, " ")

		if len(values) < 2 {
			break
		}

		values[0] = strings.TrimSpace(values[0])
		values[1] = strings.TrimSpace(values[1])

		if values[1] == "X" { // Lose
			if values[0] == "A" { // Rock
				score += 3
			} else if values[0] == "B" { // Paper
				score += 1
			} else { // Scissors
				score += 2
			}
		} else if values[1] == "Y" { // Draw
			score += 3
			if values[0] == "A" { // Rock
				score += 1
			} else if values[0] == "B" { // Paper
				score += 2
			} else { // Scissors
				score += 3
			}
		} else if values[1] == "Z" { // Win
			score += 6
			if values[0] == "A" { // Rock
				score += 2
			} else if values[0] == "B" { // Paper
				score += 3
			} else { // Scissors
				score += 1
			}
		}
	}

	return score
}
