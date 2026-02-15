package services

import (
	"os"
	"os/exec"
	"strings"

	"github.com/ukhirani/boilerplate/styles"
)

func ExecCmds(cmds []string) error {
	for i, v := range cmds {

		var cmdName string
		var cmdArgs []string

		styles.PrintStep(i+1, len(cmds), styles.CommandStyle().Render(v))

		splitCmd := strings.Split(v, " ")

		cmdName = splitCmd[0]
		if len(splitCmd) > 0 {
			cmdArgs = splitCmd[1:]
		}

		cmd := exec.Command(cmdName, cmdArgs...)

		// giving the cmd the stdiIn's and Out's
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			styles.PrintError(err.Error())
			return err
		}

	}

	return nil
}
