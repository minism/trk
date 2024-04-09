package git

import (
	"errors"
	"os"
	"os/exec"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/config"
	"github.com/spf13/cobra"
)

func MakeGitPassthroughCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git",
		Short: "Invoke a git command on the trk data repository",
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			if !checkGitExists() {
				return errors.New("unable to locate git on your system")
			}

			// TODO: Flags need to be passed through here as well.
			return invokeGitCommand(args...)
		}),
	}

	return cmd
}

func checkGitExists() bool {
	c := exec.Command("git", "version")
	return c.Run() == nil
}

func invokeGitCommand(userArg ...string) error {
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
