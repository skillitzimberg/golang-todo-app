package model

type Todo struct {
	desc string
}

func NewTodo(desc string) (todo Todo) {
	todo.desc = desc
	return
}
