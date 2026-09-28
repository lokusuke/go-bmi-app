package apperrors

import (
	"errors"
	"io"
)

var (
	ErrEOF = errors.New("入力の受付を終了しました")
	ErrUnreadable = errors.New("入力の読み取りに失敗しました")
)

func HandleInputError(err error) error {
	switch {
	case errors.Is(err, io.EOF):
		return ErrEOF
	default:
		return ErrUnreadable
	}
}