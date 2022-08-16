package main

import (
  "github.com/anyxware/todo-list"
  "github.com/anyxware/todo-list/pkg/handler"
  "github.com/anyxware/todo-list/pkg/repository"
  "github.com/anyxware/todo-list/pkg/service"
  "github.com/spf13/viper"
  "github.com/joho/godotenv"
  "github.com/sirupsen/logrus"
  "os"
)

func main() {
  logrus.SetFormatter(new(logrus.JSONFormatter))

  if err := initConfig(); err != nil {
    logrus.Fatal(err.Error())
  }

  if err := godotenv.Load(); err != nil {
    logrus.Fatal(err.Error())
  }

  config := repository.Config{
    Host:     viper.GetString("db.host"),
    Port:     viper.GetString("db.port"),
    Username: viper.GetString("db.username"),
    DBName:   viper.GetString("db.dbname"),
    SSLMode:  viper.GetString("db.sslmode"),
    Password: os.Getenv("DB_PASSWORD"),
  }

  db, err := repository.NewPostgresDB(config)
  if err != nil {
    logrus.Fatal(err.Error())
  }

  repo := repository.NewRepository(db)
  services := service.NewService(repo)
  handlers := handler.NewHandler(services)

  srv := new(todo.Server)

  if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
    logrus.Fatal(err.Error())
  }
}

func initConfig() error {
  viper.AddConfigPath("configs")
  viper.SetConfigName("config")
  return viper.ReadInConfig()
}
