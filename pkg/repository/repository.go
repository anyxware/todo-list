package repository

import (
  "github.com/jmoiron/sqlx"
  "github.com/anyxware/todo-list"
)

type Authorization interface {
  CreateUser(user todo.User) (int, error)
  GetUser(username string, password string) (todo.User, error)
}

type TodoList interface {
  Create(userId int, list todo.TodoList) (int, error)
  GetAll(userId int) ([]todo.TodoList, error)
  GetById(userId int, listId int) (todo.TodoList, error)
  Delete(userId int, listId int) error
  Update(userId int, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
  Create(listId int, item todo.TodoItems) (int, error)
  GetAll(userId int, listId int) ([]todo.TodoItems, error)
  GetById(userId int, itemId int) (todo.TodoItems, error)
  Delete(userId int, itemId int) error
  Update(userId int, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
  Authorization
  TodoList
  TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
  return &Repository{
    Authorization: NewAuthPostgres(db),
    TodoList:      NewTodoListPostgres(db),
    TodoItem:      NewTodoItemPostgres(db),
  }
}
