/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"os"
	"path/filepath"

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

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
