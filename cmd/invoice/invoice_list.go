/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package invoice

import (
	"fmt"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/pkg/model"
	"github.com/spf13/cobra"
)

var (
	flagUnpaid bool
)

func MakeInvoiceListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List invoices",
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			projects, err := shared.GetFilteredProjects()
			if err != nil {
				return err
			}

			for _, project := range projects {
				invoices, err := core.FetchInvoicesForProject(project)
				if flagUnpaid {
					invoices = model.FilterProjectInvoicesByUnpaid(invoices)
				}
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
		}),
	}

	cmd.Flags().BoolVarP(&flagUnpaid, "unpaid", "u", false, "Only show unpaid invoices")

	return cmd
}
