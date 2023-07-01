package responses

type SingUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}