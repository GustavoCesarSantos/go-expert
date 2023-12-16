package main

import (
	"github.com/GustavoCesarSantos/go-expert/rest-echo-postgresql/cmd/handlers"
	"github.com/GustavoCesarSantos/go-expert/rest-echo-postgresql/cmd/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  storage.InitDB()
  e := echo.New()
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{
      "*",
    },
    AllowHeaders: []string{
      echo.HeaderOrigin,
      echo.HeaderContentType,
      echo.HeaderAccept,
    },
  }))
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  e.GET("/", handlers.Home)
  e.POST("/users", handlers.CreateUser)
  e.Logger.Fatal(e.Start(":8080"))
}