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

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new problem-name",
	Short: "create directory for the problem, and put the template source file in it.",
	Long:  `create directory for the problem, and put the template source file in it.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("problem-name must be provided")
		}

		directory := args[0]

		templateSourceContent := GetTemplateFileContent()
		if templatePath != "" {
			content, err := os.ReadFile(templatePath)
			if err != nil {
				return fmt.Errorf("failed to read template file %s: %w", templatePath, err)
			}
			templateSourceContent = string(content)
		}

		if err := os.Mkdir(directory, 0755); err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join(directory, "main.cpp"), []byte(templateSourceContent), 0644); err != nil {
			return err
		}

		return nil
	},
}

var templatePath string

func init() {
	rootCmd.AddCommand(newCmd)
	defaultTemplate := defaultTemplatePath()
	desc := "path to template source file used for the new problem"
	if defaultTemplate != "" {
		desc = fmt.Sprintf("%s (default: %s)", desc, defaultTemplate)
	} else {
		desc = fmt.Sprintf("%s (default: $HOME/.acutils-cli/template.cpp)", desc)
	}
	newCmd.Flags().StringVar(&templatePath, "template", "", desc)
}
