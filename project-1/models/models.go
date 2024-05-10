package models

type UserData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Married   bool   `json:"married"`
}

type LoginData struct {
	Email string `json:"email"`
}

type UserEditData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Married   bool   `json:"married"`
}
