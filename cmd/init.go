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
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init contest-name",
	Short: "Initialize contest directory",
	Long: `Initialize contest directory

Create new directory for the contest. This command will also put
.vscode/settings.json on its directory.	`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New(`contest-name must be provided`)
		}
		directory := args[0]

		settingsJsonContent := GetVscodeSettingsFileContent()

		if err := os.Mkdir(directory, 0755); err != nil {
			return err
		}
		if err := os.Mkdir(filepath.Join(directory, ".vscode"), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(directory, ".vscode", "settings.json"), []byte(settingsJsonContent), 0644); err != nil {
			return err
		}
		flags := strings.Join(GetCXXFLAGS(), "\n");
		if err := os.WriteFile(filepath.Join(directory, "compile_flags.txt"), []byte(flags), 0644); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
