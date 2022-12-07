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
package day7

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Node struct {
	name     string
	size     int
	dir      bool
	children []*Node
	parent   *Node
}

func parseEntry(
	entry string,
	returnToRoot func(),
	upOneLevel func(),
	changeDir func(dir string),
	addDir func(dir string),
	addFile func(file string, size int),
) {
	split := strings.Split(entry, " ")

	if split[0] == "$" {
		if split[1] == "cd" {
			if split[2] == "/" {
				returnToRoot()
			} else if split[2] == ".." {
				upOneLevel()
			} else {
				changeDir(split[2])
			}
		}
	} else if split[0] == "dir" {
		addDir(split[1])
	} else {
		x, err := strconv.Atoi(split[0])
		cobra.CheckErr(err)

		addFile(split[1], x)
	}
}

func buildFileTree(data string) Node {
	splitData := strings.Split(data, "\n")

	root := Node{
		name:     "/",
		size:     0,
		children: []*Node{},
		dir:      true,
	}

	var current *Node = &root

	for _, v := range splitData {
		parseEntry(
			strings.TrimSpace(v),
			func() {
				current = &root
			},
			func() {
				current = current.parent
			},
			func(dir string) {
				found := false
				for _, n := range current.children {
					if n.name == dir {
						current = n
						found = true
						break
					}
				}
				if !found {
					current.children = append(current.children, &Node{
						name:     dir,
						size:     0,
						children: []*Node{},
						parent:   current,
						dir:      true,
					})
					current = current.children[len(current.children)-1]
				}
			},
			func(dir string) {
				found := false
				for _, n := range current.children {
					if n.name == dir {
						found = true
						break
					}
				}
				if !found {
					current.children = append(current.children, &Node{
						name:     dir,
						size:     0,
						children: []*Node{},
						parent:   current,
						dir:      true,
					})
				}
			},
			func(file string, size int) {
				current.children = append(current.children, &Node{
					name:     file,
					size:     size,
					children: []*Node{},
					parent:   current,
					dir:      false,
				})

				parent := current

				for ; parent != nil; parent = parent.parent {
					parent.size += size
				}
			},
		)
	}
	return root
}

func Day7Part1(data string) int {
	root := buildFileTree(data)

	queue := []*Node{&root}

	result := 0

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next.dir && next.size <= 100000 {
			result += next.size
		}

		if len(next.children) > 0 {
			queue = append(queue, next.children...)
		}
	}

	return result
}

func Day7Part2(data string) int {
	root := buildFileTree(data)

	queue := []*Node{&root}

	result := root.size

	spaceToFree := 30000000 - (70000000 - root.size)

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if next.dir && next.size >= spaceToFree && next.size < result {
			result = next.size
		}

		if len(next.children) > 0 {
			queue = append(queue, next.children...)
		}
	}

	return result
}
