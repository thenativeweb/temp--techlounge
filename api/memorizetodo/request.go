package memorizetodo

type RequestBody struct {
	Intention string `json:"intention"`
	Details   string `json:"details"`
}
