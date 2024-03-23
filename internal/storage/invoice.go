package storage

import (
	"errors"
	"os"
	"path"

	"github.com/minism/trk/pkg/model"
	"gopkg.in/yaml.v3"
)

func LoadInvoices(project model.Project) ([]model.Invoice, error) {
	data, err := os.ReadFile(project.InvoicesPath())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []model.Invoice{}, nil
		}
		return nil, err
	}

	invoices := make([]model.Invoice, 0)
	err = yaml.Unmarshal(data, &invoices)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

func SaveInvoices(project model.Project, invoices []model.Invoice) error {
	// Ensure the directory exists.
	dir := path.Dir(project.InvoicesPath())
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(invoices)
	if err != nil {
		return err
	}
	err = os.WriteFile(project.InvoicesPath(), data, 0644)
	if err != nil {
		return err
	}
	return nil
}
