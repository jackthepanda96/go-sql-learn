package controllers

import (
	"fmt"
	"todo/internal/models"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) AddTodo(id uint) (bool, error) {
	var newData models.Todo
	fmt.Print("Masukkan Aktivitas ")
	fmt.Scanln(&newData.Activity)
	newData.Owner = id
	_, err := tc.model.AddTodo(newData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (tc *TodoController) DeleteTodo(id uint) (bool, error) {
	var deleteData models.Todo
	fmt.Print("Masukkan ID Aktivitas yang akan dihapus")
	fmt.Scanln(&deleteData.ID)
	deleteData.Owner = id
	_, err := tc.model.DeleteTodo(deleteData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (tc *TodoController) FindTodo(id uint) ([]map[string]any, error) {
	data, err := tc.model.FindTodo(id)
	if err != nil {
		return nil, err
	}

	var result []map[string]any
	for _, todo := range data {
		todoMap := map[string]any{
			"id":       todo.ID,
			"activity": todo.Activity,
			"owner":    todo.Owner,
		}
		result = append(result, todoMap)
	}

	return result, nil
}

func (tc *TodoController) UpdateTodo(userId uint) {
	var todoId int
	fmt.Println("=== PERBARUI KEGIATAN ===")
	fmt.Println("Masukkan '0' untuk membatalkan pembaruan.")
	fmt.Print("Masukkan ID Kegiatan yang ingin anda tandai sebagai 'Selesai': ")
	fmt.Scanln(&todoId)
	fmt.Println()

	if todoId != 0 {
		tc.model.UpdateTodo(userId, todoId)
	}
}
