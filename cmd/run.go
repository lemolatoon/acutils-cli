/*
Copyright © 2024 lemolatoon

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
	"strings"

	"github.com/lemolatoon/acutils-cli/shell"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run problem-name",
	Short: "Compile and Run source code of specified problem-name",
	Long: `Compile and Run source code of specified problem-name

Use c++ command for compiling by default. With CXX global variable, it is used as compiler.
With CXXFLAGS in .acutils-cli.toml, you can specify compiler flags.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("problem-name must be provided")
		}

		directory := args[0]

		sourceFilePath := filepath.Join(directory, "main.cpp")

		executeFilePath := filepath.Join(directory, "a.out")
		if checkIfShouldCompile(sourceFilePath, executeFilePath) {
			flags := strings.Join(GetCXXFLAGS(), " ")
			if err := shell.Run(fmt.Sprintf("%s %s %s -o %s", GetCXX(), sourceFilePath, flags, executeFilePath)); err != nil {
				return err
			}
		}

		var executeCommand string
		if filepath.IsAbs(executeFilePath) {
			executeCommand = executeFilePath
		} else {
			executeCommand = fmt.Sprintf("./%s", executeFilePath)
		}
		if err := shell.Run(executeCommand); err != nil {
			return err
		}

		return nil
	},
}

func checkIfShouldCompile(sourceFilePath string, executeFilePath string) bool {
	executeFileInfo, err := os.Stat(executeFilePath)
	if os.IsNotExist(err) {
		return true
	}
	sourceFileInfo, err := os.Stat(sourceFilePath)
	if err != nil {
		return true
	}
	executeFileModTime := executeFileInfo.ModTime()
	sourceFileModTime := sourceFileInfo.ModTime()

	return sourceFileModTime.After(executeFileModTime)
}

func init() {
	rootCmd.AddCommand(runCmd)
}
