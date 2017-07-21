package handlers

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/user/api-template/models"
	"github.com/user/api-template/queries"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

// GetTodoByIDHandler handles request for a particular todo
func GetTodoByIDHandler(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params, db *sqlx.DB) {
	todoID, err := uuid.Parse(ps.ByName("ID"))
	if err != nil {
		ctx.Error("ID provided was not a valid UUID", fasthttp.StatusBadRequest)
		logrus.WithError(err).WithField("todoID", ps.ByName("ID")).Info("ID provided was not a valid UUID")
		return
	}

	todo, err := queries.GetTodoByID(db, todoID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			ctx.Error(fasthttp.StatusMessage(fasthttp.StatusNotFound), fasthttp.StatusNotFound)
			logrus.WithError(err).WithField("todoID", todoID).Info("No todo found with provided ID")
			return
		}
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		logrus.WithError(err).Error("Error fetching todo from database")
		return
	}

	bytes, err := json.Marshal(todo)
	if err != nil {
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		logrus.WithError(err).Error("Error marshalling data")
		return
	}

	ctx.Success("application/json", bytes)
}

// CreateTodoHandler handles request to create a new todo
func CreateTodoHandler(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params, db *sqlx.DB) {
	var todo models.Todo
	err := json.Unmarshal(ctx.PostBody(), &todo)
	if err != nil {
		ctx.Error("Unable to unmarshal post body", fasthttp.StatusBadRequest)
		logrus.WithError(err).WithField("body", ctx.PostBody()).Info("Unable to unmarshal post body")
		return
	}

	err = todo.Validate()
	if err != nil {
		ctx.Error("Invalid todo", fasthttp.StatusBadRequest)
		logrus.WithError(err).WithField("todo", todo).Info("Invalid todo")
		return
	}

	newTodo, err := queries.CreateTodo(db, &todo)
	if err != nil {
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		logrus.WithError(err).Error("Error fetching todo from database")
		return
	}

	bytes, err := json.Marshal(newTodo)
	if err != nil {
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		logrus.WithError(err).Error("Error marshalling data")
		return
	}

	ctx.Success("application/json", bytes)
}

// UpdateTodoHandler handles request to update an existing todo
func UpdateTodoHandler(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params, db *sqlx.DB) {
	var todo models.Todo
	err := json.Unmarshal(ctx.PostBody(), &todo)
	if err != nil {
		ctx.Error("Unable to unmarshal post body", fasthttp.StatusBadRequest)
		logrus.WithError(err).WithField("body", ctx.PostBody()).Info("Unable to unmarshal post body")
		return
	}

	err = todo.Validate()
	if err != nil {
		ctx.Error("Invalid todo", fasthttp.StatusBadRequest)
		logrus.WithError(err).WithField("todo", todo).Info("Invalid todo")
		return
	}

	updatedTodo, err := queries.UpdateTodo(db, &todo)
	if err != nil {
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		logrus.WithError(err).Error("Error fetching todo from database")
		return
	}

	bytes, err := json.Marshal(updatedTodo)
	if err != nil {
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		logrus.WithError(err).Error("Error marshalling data")
		return
	}

	ctx.Success("application/json", bytes)
}

func DeleteTodoHandler(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params, db *sqlx.DB) {
	todoID, err := uuid.Parse(ps.ByName("ID"))
	if err != nil {
		ctx.Error("ID provided was not a valid UUID", fasthttp.StatusBadRequest)
		logrus.WithError(err).WithField("todoID", ps.ByName("ID")).Info("ID provided was not a valid UUID")
		return
	}

	err = queries.DeleteTodo(db, todoID)
	if err != nil {
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusInternalServerError), fasthttp.StatusInternalServerError)
		logrus.WithError(err).Error("Error deleting todo from database")
		return
	}

	ctx.SuccessString("text/plain", fasthttp.StatusMessage(fasthttp.StatusOK))
}
