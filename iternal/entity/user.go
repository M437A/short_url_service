package entity

type User struct {
	Id    *uint
	Name  string
	Email string
	Links []Link
}
