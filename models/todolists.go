package models

type todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var todoList = []todo{
	{ID: 1, Title: "Todo 1", Content: "Todo 1 body"},
	{ID: 2, Title: "Todo 2", Content: "Todo 2 body"},
}

func GetAllTodos() []todo {
	return todoList
}

func GetTodoByID(id int) (*todo, error){
	for _,t := range todoList{t.ID} {
		if t.ID == id {
			return &t,nil
		}
	
	}	
	return nil, errors.New("Todo not found")}
}	

func CreateNewTodo(title,content string) (*todo, error){
	t := todo{ID: len(todoList) + 1, Title: title, Content: content}

	todoList = append(todoList,t)

	return &t, nil
}