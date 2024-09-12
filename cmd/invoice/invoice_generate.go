package invoice

import (
	"fmt"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/git"
	"github.com/minism/trk/internal/util"
	"github.com/spf13/cobra"
)

var (
	flagFutureDate string
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

				// By default we generate invoices for periods finished before today,
				// but the user can specify a future date if they really want to.
				upToDate := util.TrkToday()
				if flagFutureDate != "" {
					upToDate, err = util.ParseNaturalDate(flagFutureDate)
					if err != nil {
						return err
					}
				}

				invoices, err := core.GenerateNewInvoicesForProject(project, upToDate)
				if err != nil {
					return err
				}
				if len(invoices) < 1 {
					continue
				}
				err = git.CommitIfEnabled("Generated %d invoices for: %s\n", len(invoices), display.ColorProject(project.ID()))
				if err != nil {
					return err
				}

				display.PrintProjectInvoicesTable(invoices)
				fmt.Println()
			}

			return nil

		})}

	cmd.Flags().StringVar(&flagFutureDate, "future-date", "", "Generate invoices up to the given future date rather than using today.")

	return cmd
}
