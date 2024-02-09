package main

import (
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/gobase/api"
	"github.com/qwerty22121998/gobase/base_controller"
	"github.com/qwerty22121998/gobase/errors"
)

type ExampleController struct {
	base_controller.EchoController
}

func (e *ExampleController) Register(g *echo.Group) {
	g = g.Group("/example")
	g.GET("", e.Get)
	g.GET("/create", e.Create)
}

func (e *ExampleController) Get(c echo.Context) error {
	return base_controller.Handle(c, api.Info("Hello World"), nil)
}

func (e *ExampleController) Create(c echo.Context) error {
	return base_controller.Handle(c, nil, errors.BadRequest("Not implemented"))
}

func main() {
	e := echo.New()
	c := &ExampleController{}
	c.Register(e.Group("/api"))

	e.Start(":8080")
}
