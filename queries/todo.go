package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/user/api-template/models"
)

// GetTodoByID queries the database for a particular todo
func GetTodoByID(db *sqlx.DB, todoID uuid.UUID) (*models.Todo, error) {
	var todo models.Todo
	err := db.Get(&todo, "SELECT * FROM todo WHERE id = $1", todoID)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// CreateTodo queries the database to create a new todo
func CreateTodo(db *sqlx.DB, todo *models.Todo) (*models.Todo, error) {
	var newTodo models.Todo
	err := db.Get(&newTodo, "INSERT INTO todo (name, completed, due) VALUES ($1, $2, $3) RETURNING *", todo.Name, todo.Completed, todo.Due)
	if err != nil {
		return nil, err
	}

	return &newTodo, nil
}

// UpdateTodo queries the database to update an existing todo
func UpdateTodo(db *sqlx.DB, todo *models.Todo) (*models.Todo, error) {
	var updatedTodo models.Todo
	err := db.Get(&updatedTodo, "UPDATE todo SET name = $1, completed = $2, due = $3 WHERE id = $4 RETURNING *", todo.Name, todo.Completed, todo.Due, todo.ID)
	if err != nil {
		return nil, err
	}

	return &updatedTodo, nil
}

// DeleteTodo queries the database to delete an exisitng todo
func DeleteTodo(db *sqlx.DB, todoID uuid.UUID) error {
	_, err := db.Exec("DELETE FROM todo WHERE id = $1", todoID)
	if err != nil {
		return err
	}

	return nil
}
