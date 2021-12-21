package models

import "errors"

type Category struct {
	CategoryId   int
	Name         string
	Description  string
	ErrorExist   bool
	ErrorMessage string
}

func MaxNameCharacters(name string) error {
	if len(name) > 150 {
		return errors.New("maximum characters number for name field is 150")
	}

	return nil
}

func MaxDescriptionCharacters(name string) error {
	if len(name) > 800 {
		return errors.New("maximum characters number for description field is 150")
	}

	return nil
}
