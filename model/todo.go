package model

type Todo struct {
	Desc string
}

func NewTodo(desc string) (todo Todo) {
	todo.Desc = desc
	return
}
