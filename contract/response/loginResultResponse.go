package response

type LoginResultResponse struct {
	Data          AuthneticationResponse
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	Code          int    `json:"code"`
	Name          string `json:"username"`
	UserId        uint   `json:"id"`
	BusinessImage string `json:"businessimage"`
}
