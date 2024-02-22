package handler

import (
	"errors"
)

var (
    ErrTitleRequired       = errors.New("handler error: title is required")
    ErrTitleLength         = errors.New("handler error: limited max 10 char")
    ErrEmailRequired       = errors.New("handler error: email is required")
    ErrEmailLength         = errors.New("handler error: limited max 30 char")
    ErrInvalidEmailFormat  = errors.New("handler error: is not valid email format")
    ErrPasswordRequired    = errors.New("handler error: password is required")
    ErrPasswordLength      = errors.New("handler error: limited min 6 max 30 char")
)
