package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/fi3te/hello-go/labstack-echo/docs"
)

// @title			Echo example
// @version		1.0
// @description	A sample server that uses the Go web framework Echo.
// @contact.name	API Support
// @contact.url	https://github.com/fi3te/hello-go
// @BasePath		/
func setUpRouting(e *echo.Echo) {
	e.GET("/", home)
	e.GET("/secret", secret)

	e.GET("/users", listUsers)
	e.GET("/users/:id", getUser)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

// @Summary	Home
// @ID			home
// @Tags		General
// @Success	200	{string}	string
// @Router		/ [get]
func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// @Summary	Secret
// @ID			secret
// @Tags		General
// @Success	200	{string}	string
// @Router		/secret [get]
func secret(c echo.Context) error {
	return c.String(http.StatusOK, "This is secret.")
}

// @Summary	Lists all users
// @ID			listUsers
// @Tags		Users
// @Success	200	{object}	[]User
// @Router		/users [get]
func listUsers(c echo.Context) error {
	users := []User{
		{
			Id:       1,
			Forename: "Erika",
			Surname:  "Mustermann",
		},
		{
			Id:       2,
			Forename: "Max",
			Surname:  "Mustermann",
		},
	}
	return c.JSON(http.StatusOK, users)
}

// @Summary	Returns a specific user
// @ID			getUser
// @Tags		Users
// @Param		id	path		int	true	"Id of the user"
// @Success	200	{object}	User
// @Router		/users/{id} [get]
func getUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, User{
		Id:       id,
		Forename: "Erika",
		Surname:  "Mustermann",
	})
}
