package enums

import "errors"

const (
	Apples  = 0
	Oranges = 1
)

const (
	Summer string = "summer"
	Autumn        = "autumn"
	Winter        = "winter"
	Spring        = "spring"
)

type TypeBook uint

const (
	Adventure TypeBook = iota
	Scifi
	Action
	Terror
)

func (s TypeBook) String() string {
	switch s {
	case Adventure:
		return "Adventure"
	case Scifi:
		return "Scifi"
	case Action:
		return "Action"
	case Terror:
		return "Terror"
	}

	return "unknown"
}

func ValidateTypeBook(t TypeBook) (string, error) {
	result := t.String()
	if result == "unknown" {
		return "", errors.New("TypeBook is unknown")
	}

	return result, nil
}
