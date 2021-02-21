package models

type Fruit struct {
	Id           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Colour       string `db:"colour" json:"colour"`
	Price        string `db:"price" json:"price"`
	Quantity     string `db:"quantity" json:"quantity"`
	Descriptions string `db:"descriptions" json:"descriptions"`
}

type User struct {
	Username string `json:"username" sql:"username"`
	Email    string `json:"email" sql:"email"`
	Password string `json:"password" sql:"password"`
}

//validate:"required"
