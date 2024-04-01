package models

import "errors"

var (
	ErrNotFound   = errors.New("models: resource could not be found with the provided information")
	ErrEmailTaken = errors.New("models: email address is already in use")
)
