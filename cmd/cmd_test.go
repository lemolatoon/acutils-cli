package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func resetViperState(t *testing.T) {
	t.Helper()
	viper.Reset()
	templatePath = ""
}

func TestGetTemplateFileContentUsesDefaultTemplateFile(t *testing.T) {
	resetViperState(t)

	tmp := t.TempDir()
	t.Setenv("HOME", tmp)

	acutilsDir := filepath.Join(tmp, ".acutils-cli")
	if err := os.MkdirAll(acutilsDir, 0o755); err != nil {
		t.Fatalf("failed to create default dir: %v", err)
	}
	want := "# default template\n"
	if err := os.WriteFile(filepath.Join(acutilsDir, "template.cpp"), []byte(want), 0o644); err != nil {
		t.Fatalf("failed to write default template: %v", err)
	}

	got := GetTemplateFileContent()
	if got != want {
		t.Fatalf("template content mismatch:\nwant: %q\ngot : %q", want, got)
	}
}

func TestGetTemplateFileContentUsesExplicitTemplatePath(t *testing.T) {
	resetViperState(t)

	tmp := t.TempDir()
	templateFile := filepath.Join(tmp, "custom.cpp")
	want := "// custom template\n"
	if err := os.WriteFile(templateFile, []byte(want), 0o644); err != nil {
		t.Fatalf("failed to write custom template: %v", err)
	}

	viper.Set(TEMPLATE_FILE_KEY, templateFile)

	if got := GetTemplateFileContent(); got != want {
		t.Fatalf("template content mismatch:\nwant: %q\ngot : %q", want, got)
	}
}

func TestInitCmdCreatesContestScaffolding(t *testing.T) {
	resetViperState(t)

	tmp := t.TempDir()
	origWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %v", err)
	}
	defer func() {
		_ = os.Chdir(origWD)
	}()
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}

	if err := initCmd.RunE(initCmd, []string{"abc001"}); err != nil {
		t.Fatalf("init command failed: %v", err)
	}

	contestDir := filepath.Join(tmp, "abc001")
	if _, err := os.Stat(contestDir); err != nil {
		t.Fatalf("contest dir not created: %v", err)
	}

	settingsPath := filepath.Join(contestDir, ".vscode", "settings.json")
	settings, err := os.ReadFile(settingsPath)
	if err != nil {
		t.Fatalf("settings.json missing: %v", err)
	}
	if string(settings) != VSCODE_TEMPLATE_SETTINGS_DEFAULT {
		t.Fatalf("settings.json content mismatch")
	}

	flagsPath := filepath.Join(contestDir, "compile_flags.txt")
	flags, err := os.ReadFile(flagsPath)
	if err != nil {
		t.Fatalf("compile_flags.txt missing: %v", err)
	}
	if string(flags) != strings.Join(DEFAULT_CXXFLAGS, "\n") {
		t.Fatalf("compile flags mismatch: %q", string(flags))
	}
}

func TestNewCmdCreatesProblemWithTemplate(t *testing.T) {
	resetViperState(t)

	tmp := t.TempDir()
	t.Setenv("HOME", tmp)

	templateFile := filepath.Join(tmp, "my-template.cpp")
	want := "// from flag\n"
	if err := os.WriteFile(templateFile, []byte(want), 0o644); err != nil {
		t.Fatalf("failed to write template: %v", err)
	}

	origWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %v", err)
	}
	defer func() {
		_ = os.Chdir(origWD)
	}()
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}

	templatePath = templateFile
	if err := newCmd.RunE(newCmd, []string{"a"}); err != nil {
		t.Fatalf("new command failed: %v", err)
	}

	content, err := os.ReadFile(filepath.Join(tmp, "a", "main.cpp"))
	if err != nil {
		t.Fatalf("main.cpp missing: %v", err)
	}
	if string(content) != want {
		t.Fatalf("main.cpp content mismatch:\nwant: %q\ngot : %q", want, string(content))
	}
}

func TestCheckIfShouldCompile(t *testing.T) {
	tmp := t.TempDir()
	source := filepath.Join(tmp, "main.cpp")
	execFile := filepath.Join(tmp, "a.out")

	if err := os.WriteFile(source, []byte("code"), 0o644); err != nil {
		t.Fatalf("failed to write source: %v", err)
	}

	if !checkIfShouldCompile(source, execFile) {
		t.Fatalf("expected compilation when executable is missing")
	}

	if err := os.WriteFile(execFile, []byte{}, 0o755); err != nil {
		t.Fatalf("failed to write executable: %v", err)
	}

	newer := time.Now().Add(2 * time.Second)
	if err := os.Chtimes(source, newer, newer); err != nil {
		t.Fatalf("failed to touch source: %v", err)
	}
	if !checkIfShouldCompile(source, execFile) {
		t.Fatalf("expected compilation when source is newer")
	}

	older := time.Now().Add(-2 * time.Second)
	if err := os.Chtimes(source, older, older); err != nil {
		t.Fatalf("failed to touch source: %v", err)
	}
	future := time.Now().Add(5 * time.Second)
	if err := os.Chtimes(execFile, future, future); err != nil {
		t.Fatalf("failed to touch exec: %v", err)
	}

	if checkIfShouldCompile(source, execFile) {
		t.Fatalf("did not expect compilation when executable is newer")
	}
}

func TestClipFallsBackWithoutClipboardCommand(t *testing.T) {
	resetViperState(t)

	tmp := t.TempDir()
	contestDir := filepath.Join(tmp, "p1")
	if err := os.Mkdir(contestDir, 0o755); err != nil {
		t.Fatalf("failed to create problem dir: %v", err)
	}
	want := "// answer\n"
	if err := os.WriteFile(filepath.Join(contestDir, "main.cpp"), []byte(want), 0o644); err != nil {
		t.Fatalf("failed to write source: %v", err)
	}

	origWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %v", err)
	}
	defer func() {
		_ = os.Chdir(origWD)
	}()
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}

	// Ensure clip.exe / pbcopy are not found.
	t.Setenv("PATH", filepath.Join(tmp, "empty-bin"))

	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to pipe stdout: %v", err)
	}
	os.Stdout = w
	defer func() {
		_ = w.Close()
		os.Stdout = origStdout
	}()

	if err := clip("p1"); err != nil {
		t.Fatalf("clip failed: %v", err)
	}

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close writer: %v", err)
	}
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}

	if !bytes.Contains(out, []byte(want)) {
		t.Fatalf("expected fallback output to include source content")
	}
}
