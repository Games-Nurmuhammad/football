package methods

import (
	"os"
	"os/exec"
)

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
