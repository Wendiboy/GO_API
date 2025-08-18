package userService

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
