package pendingtodos

type ResponseBody struct {
	ToDos []ResponseBodyToDo `json:"todos"`
}

type ResponseBodyToDo struct {
	ID        string `json:"id"`
	Intention string `json:"intention"`
	Details   string `json:"details"`
}
