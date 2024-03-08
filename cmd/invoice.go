/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/model"
	"github.com/spf13/cobra"
)

func runListInvoicesCmd(cmd *cobra.Command, args []string) error {
	projects, err := core.GetAllProjects()
	if err != nil {
		return err
	}

	// Optionally filter by a single project.
	if flagProject != "" {
		project, err := core.FilterProjectsByIdFuzzy(projects, flagProject)
		if err != nil {
			return err
		}
		projects = []model.Project{project}
	}

	for _, project := range projects {
		invoices, err := core.FetchInvoicesForProject(project)
		if err != nil {
			return err
		}
		log.Printf("Project: %s\n", display.ColorProject(project.ID()))
		display.PrintInvoicesTable(invoices)
		fmt.Println()
	}

	return nil
}

func runGenerateInvoiceCmd(cmd *cobra.Command, args []string) error {
	projects, err := core.GetAllProjects()
	if err != nil {
		return err
	}

	for _, project := range projects {
		invoices, err := core.GenerateInvoicesForProject(project)
		if err != nil {
			return err
		}
		log.Printf("Project: %s\n", display.ColorProject(project.ID()))
		display.PrintInvoicesTable(invoices)
		fmt.Println()
	}

	return nil
}

var invoicesCmd = &cobra.Command{
	Use:     "invoice",
	Aliases: []string{"invoices"},
	Short:   "View and manage invoices",
}

var listInvoicesCmd = &cobra.Command{
	Use:   "list",
	Short: "List invoices",
	Run:   wrapCommand(runListInvoicesCmd),
}

var generateInvoicesCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generate invoices",
	Run:   wrapCommand(runGenerateInvoiceCmd),
}

func init() {
	rootCmd.AddCommand(invoicesCmd)
	invoicesCmd.AddCommand(listInvoicesCmd)
	invoicesCmd.AddCommand(generateInvoicesCommand)
}
