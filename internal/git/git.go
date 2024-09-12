package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/internal/storage"
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

func CommitIfEnabled(format string, args ...any) error {
	cfg, err := storage.LoadConfig()
	if err != nil {
		return err
	}
	msg := fmt.Sprintf(format, args...)
	if cfg.AutoCommit {
		InvokeGitCommand("commit", "-am", msg)
	} else {
		log.Print(msg)
	}
	return nil
}
