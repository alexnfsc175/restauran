package models

// Todo Estrutura da tabela
type Todo struct { //  application/json ou application/x-www-form-urlencoded
	ID          uint   `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}

// TableName Retorna o nome da tabela
func (b *Todo) TableName() string {
	return "todo"
}
