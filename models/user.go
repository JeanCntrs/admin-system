package models

type User struct {
	UserId       int
	Username     string
	Password     string
	PersonId     int
	RoleTypeId   int
	RoleTypeName string
	Fullname     string
	Errors       []string
}
