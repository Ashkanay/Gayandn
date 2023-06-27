package request

type LoginRequest struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
