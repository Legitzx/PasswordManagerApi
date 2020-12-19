package model

type User struct {
	Id       string    `json:"id"`
	Email    string    `json:"email"`
	Iv       string    `json:"iv"`
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseResult struct {
	Result string `json:"result"`
	Token  string `json:"token"`
}
