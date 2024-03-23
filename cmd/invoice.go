/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/pkg/model"
	"github.com/spf13/cobra"
)

func runListInvoicesCmd(cmd *cobra.Command, args []string) error {
	projects, err := getFilteredProjects()
	if err != nil {
		return err
	}

	for _, project := range projects {
		invoices, err := core.FetchInvoicesForProject(project)
		if err != nil {
			return err
		}
		if len(invoices) < 1 {
			continue
		}
		fmt.Printf("Project: %s\n", display.ColorProject(project.ID()))
		display.PrintProjectInvoicesTable(invoices)
		fmt.Println()
	}

	return nil
}

func runGenerateInvoiceCmd(cmd *cobra.Command, args []string) error {
	projects, err := getFilteredProjects()
	if err != nil {
		return err
	}

	for _, project := range projects {
		invoices, err := core.GenerateNewInvoicesForProject(project)
		if err != nil {
			return err
		}
		if len(invoices) < 1 {
			continue
		}
		fmt.Printf("Generated %d invoices for: %s\n", len(invoices), display.ColorProject(project.ID()))
		display.PrintProjectInvoicesTable(invoices)
		fmt.Println()
	}

	return nil
}

func runDeleteInvoiceCmd(cmd *cobra.Command, args []string) error {
	id := model.ProjectInvoiceId(args[0])
	err := core.DeleteProjectInvoiceById(id)
	// TODO: Rethink what should go to stderr and what should go to stdout for cases like this, use git as a guiding example.
	fmt.Printf("Deleted invoice %s\n", id)
	return err
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

var deleteInvoiceCommand = &cobra.Command{
	Use:   "delete <invoice_id>",
	Short: "Delete invoices",
	Run:   wrapCommand(runDeleteInvoiceCmd),
}

func init() {
	rootCmd.AddCommand(invoicesCmd)
	invoicesCmd.AddCommand(listInvoicesCmd)
	invoicesCmd.AddCommand(generateInvoicesCommand)
	invoicesCmd.AddCommand(deleteInvoiceCommand)

	deleteInvoiceCommand.Args = cobra.ExactArgs(1)
}
