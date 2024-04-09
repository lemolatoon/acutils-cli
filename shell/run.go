package shell

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(command string) error {
	fmt.Printf("+%s\n", command)
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
