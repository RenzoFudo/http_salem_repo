package storage

import (
	"errors"

	errtext
)

var ErrInvalidAuthData = errors.New(errtext.InvalidAuthDataError)
var ErrUserNotFound = errors.New(errtext.UserNotFoundError)
var ErrBookNotFound = errors.New(errtext.BookNotFoundError)
var ErrBooksListEmpty = errors.New(errtext.BooksListEmptyError)
