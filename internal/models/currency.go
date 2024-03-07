package models

type Currency struct {
	ID   uint   `sql:"id"`
	Name string `sql:"name"`
}
