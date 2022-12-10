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
package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Point struct {
	x int
	y int
}

func (p *Point) toString() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

func (p *Point) isAdjacent(p2 *Point) bool {
	return (p.x-1 == p2.x && p.y-1 == p2.y) || (p.x-1 == p2.x && p.y == p2.y) || (p.x-1 == p2.x && p.y+1 == p2.y) || (p.x == p2.x && p.y-1 == p2.y) || (p.x == p2.x && p.y == p2.y) || (p.x == p2.x && p.y+1 == p2.y) || (p.x+1 == p2.x && p.y-1 == p2.y) || (p.x+1 == p2.x && p.y == p2.y) || (p.x+1 == p2.x && p.y+1 == p2.y)
}

func (p *Point) moveTo(p2 *Point) {
	if p.x > p2.x {
		p.x -= 1
	} else if p.x < p2.x {
		p.x += 1
	}

	if p.y > p2.y {
		p.y -= 1
	} else if p.y < p2.y {
		p.y += 1
	}
}

func Day9Part1(data string) int {
	splitData := strings.Split(data, "\n")

	head, tail := Point{}, Point{}
	tailVisited := map[string]bool{tail.toString(): true}

	for _, v := range splitData {
		val := strings.TrimSpace(v)
		instructions := strings.Split(val, " ")

		amount, err := strconv.Atoi(instructions[1])
		cobra.CheckErr(err)

		for i := 0; i < amount; i++ {
			if instructions[0] == "R" {
				head.x += 1
			} else if instructions[0] == "U" {
				head.y += 1
			} else if instructions[0] == "L" {
				head.x -= 1
			} else if instructions[0] == "D" {
				head.y -= 1
			}

			if !tail.isAdjacent(&head) {
				tail.moveTo(&head)
			}

			tailVisited[tail.toString()] = true
		}
	}

	return len(tailVisited)
}

func Day9Part2(data string) int {
	splitData := strings.Split(data, "\n")

	rope := [10]Point{}
	tailVisited := map[string]bool{rope[9].toString(): true}

	for _, v := range splitData {
		val := strings.TrimSpace(v)
		instructions := strings.Split(val, " ")

		amount, err := strconv.Atoi(instructions[1])
		cobra.CheckErr(err)

		for i := 0; i < amount; i++ {
			if instructions[0] == "R" {
				rope[0].x += 1
			} else if instructions[0] == "U" {
				rope[0].y += 1
			} else if instructions[0] == "L" {
				rope[0].x -= 1
			} else if instructions[0] == "D" {
				rope[0].y -= 1
			}

			for j := 1; j < len(rope); j++ {
				if !rope[j].isAdjacent(&rope[j-1]) {
					rope[j].moveTo(&rope[j-1])
				}
			}

			tailVisited[rope[9].toString()] = true
		}
	}

	return len(tailVisited)
}
