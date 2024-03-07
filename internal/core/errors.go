package core

import "errors"

var (
	ErrProjectNotFound         error = errors.New("project not found")
	ErrMultipleProjectsMatched error = errors.New("multiple projects were matched")
	ErrHoursExceedsLimit       error = errors.New("exceeds total hours in a day")
)
