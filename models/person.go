package models

import "time"

type Person struct {
	PersonId          int
	Name              string
	FatherLastName    string
	MotherLastName    string
	Fullname          string
	Birthday          time.Time
	FormattedBirthday string
	NameTypePerson    string
}
