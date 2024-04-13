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
	"fmt"
	"os"
	"path/filepath"

	"github.com/hairyhenderson/go-which"
	"github.com/lemolatoon/acutils-cli/shell"
	"github.com/spf13/cobra"
)

// clipCmd represents the clip command
var clipCmd = &cobra.Command{
	Use:   "clip problem-name",
	Short: "Copy the source code to the clipboard.",
	Long:  `Clip subcommand copies the source code in specified folder to the clipboard.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("problem-name must be provided")
		}

		problemName := args[0]
		return clip(problemName)
	},
}

func clip(problemName string) error {
	sourceFilePath := filepath.Join(problemName, "main.cpp")

	if which.Found("clip.exe") {
		if err := shell.Run(fmt.Sprintf("clip.exe < %s", sourceFilePath)); err != nil {
			return err
		}
		return nil
	}

	if which.Found("pbcopy") {
		if err := shell.Run(fmt.Sprintf("pbcopy < %s", sourceFilePath)); err != nil {
			return err
		}
		return nil
	}

	content, err := os.ReadFile(sourceFilePath)
	if err != nil {
		return err
	}
	fmt.Printf(`we cannot find the way to copy to the clipboard.
Please copy the source code manually.
################################################################
%s
################################################################`, content)

	return nil
}

func init() {
	rootCmd.AddCommand(clipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
