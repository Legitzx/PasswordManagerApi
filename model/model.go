package model

type User struct {
	VaultKey string    `json:"vaultKey"`
	Email    string    `json:"email"`
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
