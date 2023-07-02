package responses

type SingUp struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Login struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}
