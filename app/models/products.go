package models

type Product struct {
	Id	string	`json:"id"`
	Title	string	`json:"title"`
	Price	int	`json:"price"`
}