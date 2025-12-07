package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecCmds(cmds []string) error {
	for _, v := range cmds {

		var cmdName string
		var cmdArgs []string

		fmt.Println("[EXEC]", v)

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
			fmt.Println(" - [ERR]", err)
			return err
		}

	}

	return nil
}
