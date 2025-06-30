package repository

import "errors"

var (
	ErrDuplicateIruyanID = errors.New("iruyan_id is already taken")
	ErrDuplicateEmail    = errors.New("email is already registered")
)
