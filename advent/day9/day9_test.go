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
package day9_test

import (
	"testing"

	"github.com/Sloff/advent-of-code-2022/advent/day9"
)

var data = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestDay9Part1(t *testing.T) {
	expectedResult := 13
	result := day9.Day9Part1(data)

	if result != expectedResult {
		t.Logf("should be %v, but got %v\n", expectedResult, result)
		t.Fail()
	}
}

var data2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestDay9Part2(t *testing.T) {
	expectedResult := 36
	result := day9.Day9Part2(data2)

	if result != expectedResult {
		t.Logf("should be %v, but got %v\n", expectedResult, result)
		t.Fail()
	}
}
