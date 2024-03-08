package config

import (
	"fmt"
	"os"
	"path"
)

func GetConfigPath() string {
	return path.Join(GetUserAppDir(), "config.yaml")
}

func GetWorkLogDir() string {
	return path.Join(GetUserAppDir(), "worklog")
}

func GetInvoicesDir() string {
	return path.Join(GetUserAppDir(), "invoices")
}

func GetProjectWorkLogPath(projectId string) string {
	return path.Join(GetWorkLogDir(), fmt.Sprintf("%s.csv", projectId))
}

func GetProjectInvoicesPath(projectId string) string {
	return path.Join(GetInvoicesDir(), fmt.Sprintf("%s_invoices.yaml", projectId))
}

// The directory all trk's application data is located for the current user.
func GetUserAppDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(home, ".trk")
}
