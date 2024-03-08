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

var invoicesCmd = &cobra.Command{
	Use:     "invoice",
	Aliases: []string{"invoices"},
	Short:   "View and manage invoices",
}

var listInvoicesCmd = &cobra.Command{
	Use:   "list",
	Short: "List invoices",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := core.GetAllProjects()
		if err != nil {
			// TODO: Command runners should just have this logic wrapped.
			log.Fatal(err)
		}

		// Optionally filter by a single project.
		if flagProject != "" {
			project, err := core.FilterProjectsByIdFuzzy(projects, flagProject)
			if err != nil {
				// TODO: Share the error handling which dumps project IDs here.
				log.Fatal(err)
			}
			projects = []model.Project{project}
		}

		for _, project := range projects {
			invoices, err := core.FetchInvoicesForProject(project)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Project: %s\n", display.ColorProject(project.ID()))
			display.PrintInvoicesTable(invoices)
			fmt.Println()
		}
	},
}

var generateInvoicesCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generate invoices",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := core.GetAllProjects()
		if err != nil {
			log.Fatal(err)
		}

		for _, project := range projects {
			invoices, err := core.GenerateInvoicesForProject(project)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Project: %s\n", display.ColorProject(project.ID()))
			display.PrintInvoicesTable(invoices)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(invoicesCmd)
	invoicesCmd.AddCommand(listInvoicesCmd)
	invoicesCmd.AddCommand(generateInvoicesCommand)
}
