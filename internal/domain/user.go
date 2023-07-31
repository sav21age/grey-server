package domain

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Fullname  string `json:"fullname" binding:"required"`
	Age       int    `json:"age" binding:"required,gt=18"`
	IsMarried bool   `json:"is_married" binding:"required,eq=false"`
	Password  string `json:"password" binding:"required,min=8"`
}

type UserSignUpInput struct {
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Age       int    `json:"age" binding:"required,numeric,gte=18"`
	IsMarried bool   `json:"is_married"`
	Password  string `json:"password" binding:"required,min=8"`
}
