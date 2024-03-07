package core

import "errors"

var (
	ErrProjectNotFound         error = errors.New("project not found")
	ErrMultipleProjectsMatched error = errors.New("multiple projects were matched")
)
