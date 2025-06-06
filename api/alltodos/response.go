package alltodos

type ResponseBody struct {
	ToDos []ResponseBodyToDo `json:"todos"`
}

type ResponseBodyToDo struct {
	ID          string `json:"id"`
	Intention   string `json:"intention"`
	IsCompleted bool   `json:"isCompleted"`
}
