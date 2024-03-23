package core

import "errors"

var (
	ErrProjectNotFound         error = errors.New("project not found")
	ErrMultipleProjectsMatched error = errors.New("multiple projects were matched")
	ErrInvoiceNotFound         error = errors.New("invoice not found")
	ErrHoursExceedsLimit       error = errors.New("exceeds total hours in a day")
	ErrNotImplemented          error = errors.New("not yet implemented")
)
