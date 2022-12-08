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
package day8_test

import (
	"testing"

	"github.com/Sloff/advent-of-code-2022/advent/day8"
)

var data = `30373
25512
65332
33549
35390`

func TestDay8Part1(t *testing.T) {
	expectedResult := 21
	result := day8.Day8Part1(data)

	if result != expectedResult {
		t.Logf("should be %v, but got %v\n", expectedResult, result)
		t.Fail()
	}
}

func TestDay8Part2(t *testing.T) {
	expectedResult := 8
	result := day8.Day8Part2(data)

	if result != expectedResult {
		t.Logf("should be %v, but got %v\n", expectedResult, result)
		t.Fail()
	}
}
