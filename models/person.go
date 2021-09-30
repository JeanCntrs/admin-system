package models

type Person struct {
	PersonID int
	Names    string
	Surnames string
	Age      int
	IsMan    bool
	Friends  []string
	Hobbies  []Hobby
}
