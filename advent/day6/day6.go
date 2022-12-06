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
package day6

import "strings"

func Day6Part1(data string) int {
	i := 3
	for ; i < len(data); i++ {
		contains1 := strings.Contains(data[i-2:i], string(data[i-3]))
		contains2 := strings.Contains(string(data[i-3])+data[i-1:i], string(data[i-2]))
		contains3 := strings.Contains(data[i-3:i-2]+string(data[i]), string(data[i-1]))
		contains4 := strings.Contains(data[i-3:i-1], string(data[i]))

		if !contains1 && !contains2 && !contains3 && !contains4 {
			return i + 1
		}
	}

	return i + 1
}

func Day6Part2(data string) int {
	amount := 13
	i := amount
	for ; i < len(data); i++ {
		double := false
		for x := i - amount; x < i; x++ {
			if strings.Count(data[i-amount:i+1], string(data[x])) > 1 {
				double = true
				break
			}
		}

		if !double {
			break
		}
	}

	return i + 1
}
