package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Activity string
	Mark     bool
	Owner    uint
}

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) *TodoModel {
	return &TodoModel{
		db: connection,
	}
}

func (tm *TodoModel) AddTodo(newData Todo) (Todo, error) {
	newData.Mark = false
	err := tm.db.Create(&newData).Error
	if err != nil {
		return Todo{}, err
	}

	return newData, nil
}

func (tm *TodoModel) DeleteTodo(deleteData Todo) (Todo, error) {
	// deleteData.Mark = true

	// query := ` UPDATE "be24"."todos" SET "deleted_at"= ?
	// WHERE (owner = ? AND activity = ?) AND "todos"."deleted_at" IS NULL `
	// err := tm.db.Exec(query, &deleteData.UpdatedAt, &deleteData.Owner, &deleteData.Activity).Error
	query := tm.db.Delete(&deleteData)
	// fmt.Println(query)
	if query.Error != nil {
		return Todo{}, query.Error
	}

	if query.RowsAffected < 1 {
		return Todo{}, gorm.ErrRecordNotFound
	}

	return deleteData, nil
}

func (tm *TodoModel) FindTodo(owner uint) ([]Todo, error) {
	var todo []Todo
	err := tm.db.Where("owner = ?", owner).Find(&todo).Error
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (tm *TodoModel) UpdateTodo(userID uint, todoID int) error {
	err := tm.db.Model(&Todo{}).Where("id = ? AND owner = ?", todoID, userID).Update("mark", true).Error
	return err
}
