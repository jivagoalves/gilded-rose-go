package name

import "strings"

type name string

type N interface {
	Value() string
	private()
}

type ErrBlankName struct {
}

func (e ErrBlankName) Error() string {
	return "Name can't be blank."
}

func (n name) Value() string {
	return string(n)
}

func (n name) private() {
	// nothing
}

func New(s string) (N, error) {
	if strings.Trim(s, " ") == "" {
		return nil, ErrBlankName{}
	}
	return name(s), nil
}
