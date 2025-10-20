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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "acutils-cli",
	Version: "0.1.0",
	Short:   "acutils is a Atcoder utilities CLI developed only for lemolatoon.",
	Long: `acutils is a Atcoder utilities CLI developed only for lemolatoon.

This application is a tool to generate template directory/files for contests 
to quickly start solving problems of the contest.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.acutils-cli/config.toml)")
}

const TEMPLATE_FILE_KEY = "TEMPLATE_FILE"
const TEMPLATE_DEFAULT = `#include <bits/stdc++.h>
using namespace std;

int main() {
  cout << "Hello!\n";
  return 0;
}
`

func configDir() string {
	configPath := viper.ConfigFileUsed()
	if configPath == "" {
		return ""
	}

	return filepath.Dir(configPath)
}

func defaultTemplatePath() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return ""
	}

	return filepath.Join(home, ".acutils-cli", "template.cpp")
}

func GetTemplateFileContent() string {
	templateFilepath := viper.GetString(TEMPLATE_FILE_KEY)
	if templateFilepath == "" {
		if defaultPath := defaultTemplatePath(); defaultPath != "" {
			if content, err := os.ReadFile(defaultPath); err == nil {
				return string(content)
			}
		}
		return TEMPLATE_DEFAULT
	}
	templateFullpath := templateFilepath
	if !filepath.IsAbs(templateFilepath) {
		if dir := configDir(); dir != "" {
			templateFullpath = filepath.Join(dir, templateFilepath)
		}
	}

	content, err := os.ReadFile(templateFullpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read template file: %s\n", templateFullpath)
		if defaultPath := defaultTemplatePath(); defaultPath != "" {
			if content, err := os.ReadFile(defaultPath); err == nil {
				return string(content)
			}
		}
		return TEMPLATE_DEFAULT
	}
	return string(content)
}

func GetCXX() string {
	cxx := viper.GetString("CXX")
	if cxx != "" {
		return cxx
	}
	cxx = os.Getenv("CXX")
	if cxx != "" {
		return cxx
	}

	return "c++"
}

const DEFAULT_CXXFLAGS = "-g -Wall -Wextra -fsanitize=undefined,address -std=c++20"

func GetCXXFLAGS() string {
	cxxflags := viper.GetString("CXXFLAGS")
	if cxxflags != "" {
		return cxxflags
	}

	return DEFAULT_CXXFLAGS
}

const VSCODE_TEMPLATE_SETTINGS_FILE_KEY = "VSCODE_TEMPLATE_SETTINGS_FILE"
const VSCODE_TEMPLATE_SETTINGS_DEFAULT = `
{
        "[cpp]": {
                "editor.defaultFormatter": "xaver.clang-format"
        },
        "clang-format.executable": "/usr/bin/clang-format-14",
        "clangd.fallbackFlags": [
                "-std=c++20"
        ],
        "editor.formatOnSave": true,
        "github.copilot.enable": {
                "*": false,
                "plaintext": false,
                "markdown": false,
                "scminput": false
        }
}
`

func GetVscodeSettingsFileContent() string {
	settingsJsonPath := viper.GetString(VSCODE_TEMPLATE_SETTINGS_FILE_KEY)
	if settingsJsonPath == "" {
		return VSCODE_TEMPLATE_SETTINGS_DEFAULT
	}
	settingsJsonFullpath := filepath.Join(configDir(), settingsJsonPath)

	content, err := os.ReadFile(settingsJsonFullpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read settings.json file: %s\n", settingsJsonFullpath)
		return VSCODE_TEMPLATE_SETTINGS_DEFAULT
	}
	return string(content)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".acutils-cli" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".acutils-cli"))
		viper.SetConfigType("toml")
		viper.SetConfigName("config.toml")
	}

	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}
