package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
)

func main(){
  //Create view engine
  viewsEngine := html.New("./views", ".html")

  //Start new fiber instance
  app := fiber.New(fiber.Config{
    Views: viewsEngine,
  })

  //Statiic route and directory
  app.Static("/static/", "./static")

  //Create a "ping" handler to test the server
  app.Get("/ping", func(ctx *fiber.Ctx) error{
    return ctx.SendString("Welcome to fiber.")
  })

  //Start the http server
  app.Listen(":3000")
}
