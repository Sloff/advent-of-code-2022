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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	tpl "github.com/Sloff/advent-of-code-2022/template"

	"github.com/spf13/cobra"
)

// newDayCmd represents the new command
var newDayCmd = &cobra.Command{
	Use:          "new",
	Short:        "Create a new advent day",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		day := args[0]
		fullDay := fmt.Sprintf("day%s", day)

		d := struct {
			Day string
		}{
			Day: day,
		}

		// Advent Dir
		adventPath := filepath.Join("advent", fullDay)

		if _, err := os.Stat(adventPath); !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%s already exists", adventPath)
		}

		err := os.MkdirAll(adventPath, os.ModePerm)
		cobra.CheckErr(err)

		// Test file
		testFile, err := os.Create(filepath.Join(adventPath, fmt.Sprintf("%s_test.go", fullDay)))
		cobra.CheckErr(err)
		defer testFile.Close()

		testTemplate := template.Must(template.New("test").Parse(tpl.AdventDayTestTemplate()))
		err = testTemplate.Execute(testFile, d)
		cobra.CheckErr(err)

		fmt.Println(testFile.Name(), "created")

		// Advent file
		adventFile, err := os.Create(fmt.Sprintf("%s/%s.go", adventPath, fullDay))
		cobra.CheckErr(err)
		defer adventFile.Close()

		adventTemplate := template.Must(template.New("advent").Parse(tpl.AdventDayTemplate()))
		err = adventTemplate.Execute(adventFile, d)
		cobra.CheckErr(err)

		fmt.Println(adventFile.Name(), "created")

		// Advent Dir
		cmdFilePath := filepath.Join("cmd", fmt.Sprintf("%s.go", fullDay))

		if _, err := os.Stat(cmdFilePath); !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%s already exists", cmdFilePath)
		}

		// Cmd file
		cmdFile, err := os.Create(cmdFilePath)
		cobra.CheckErr(err)
		defer cmdFile.Close()

		cmdTemplate := template.Must(template.New("cmd").Parse(tpl.CommandTemplate()))
		err = cmdTemplate.Execute(cmdFile, d)
		cobra.CheckErr(err)

		fmt.Println(cmdFile.Name(), "created")

		// Input file
		inputFile, err := os.Create(filepath.Join("resources", fmt.Sprintf("%s.txt", fullDay)))
		cobra.CheckErr(err)
		defer inputFile.Close()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(newDayCmd)
}
