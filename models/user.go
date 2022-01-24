package models

type User struct {
	UserId       int
	Username     string
	Password     string
	PersonId     int
	RoleTypeId   int
	RoleTypeName string
	PersonName   string
	Errors       []string
}
