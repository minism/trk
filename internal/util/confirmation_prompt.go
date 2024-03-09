package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskUserForConfirmation(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s [y/N]: ", prompt)
		response, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		response = strings.ToLower(strings.TrimSpace(response))
		return response == "y" || response == "yes"
	}
}
