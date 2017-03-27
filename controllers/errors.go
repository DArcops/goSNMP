package controllers

import "errors"

var (
	//ErrDuplicate is used when exists a repeated unique value.
	ErrDuplicate = errors.New("duplicated unique value")
	//ErrNotFound is used when can't find a resource, for example, an agent.
	ErrNotFound = errors.New("unable to find resource")
)
