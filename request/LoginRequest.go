package request

type LoginRequest struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"pwd"`
	Value    string `json:"value"`
}
