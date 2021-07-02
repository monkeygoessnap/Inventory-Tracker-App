package database

import "errors"

//error variables
var (
	ErrUserTaken = errors.New("username taken")
	ErrInternal  = errors.New("internal error")
	ErrNotFound  = errors.New("not found")
)
