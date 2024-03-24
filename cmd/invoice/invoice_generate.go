package invoice

import (
	"fmt"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/spf13/cobra"
)

func MakeInvoiceGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate invoices",
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {

			projects, err := shared.GetFilteredProjects()
			if err != nil {
				return err
			}

			for _, project := range projects {
				if project.InvoiceInterval != config.InvoiceIntervalBimonthly {
					fmt.Printf("Skipping %s because it uses invoice interval %s\n", display.ColorProject(project.Name), project.InvoiceInterval)
					continue
				}

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

		})}

	return cmd
}
