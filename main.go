package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/safaaleigh/api-template/app"
	"github.com/safaaleigh/api-template/handlers"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

// ConnectedHandler defines a handler with accepts a database connection as an argument
type ConnectedHandler func(*fasthttp.RequestCtx, fasthttprouter.Params, *sqlx.DB)

// ConnectDatabase defines a middle to pass a database connection to a handler
func ConnectDatabase(h ConnectedHandler, db *sqlx.DB) fasthttprouter.Handle {
	return fasthttprouter.Handle(func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
		h(ctx, ps, db)
		return
	})
}

func main() {
	cfg := app.LoadConfig()

	logger := logrus.New()

	db := sqlx.MustConnect("postgres", "postgres://postgres:"+cfg.DB.User+"@127.0.0.1:5432/"+cfg.DB.Name+"?sslmode=disable")

	router := fasthttprouter.New()
	router.GET("/todos/:ID", ConnectDatabase(handlers.GetTodoByIDHandler, db))
	router.POST("/todos", ConnectDatabase(handlers.CreateTodoHandler, db))
	router.PUT("/todos", ConnectDatabase(handlers.UpdateTodoHandler, db))
	router.DELETE("/todos/:ID", ConnectDatabase(handlers.DeleteTodoHandler, db))

	address := fmt.Sprintf(":%v", cfg.Server.Port)
	logger.Infof("Server is started at %v", address)
	panic(fasthttp.ListenAndServe(address, router.Handler))
}
