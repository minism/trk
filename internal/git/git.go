package git

import (
	"os"
	"os/exec"

	"github.com/minism/trk/internal/config"
)

func CheckGitExists() bool {
	c := exec.Command("git", "version")
	return c.Run() == nil
}

func InvokeGitCommand(userArg ...string) error {
	arg := []string{
		"-C",
		config.GetUserAppDir(),
	}
	arg = append(arg, userArg...)
	c := exec.Command("git", arg...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
