package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var Users *userRouting

type userRouting struct {
}
type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

func (t *userRouting) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func (t *userRouting) Add(c echo.Context) error {
	u := &user{
		ID:   seq,
		Name: "italo",
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)

}

func (t *userRouting) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)

}
func (t *userRouting) Update(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

var Get *getRouting

type getRouting struct {
}

func (t *getRouting) Novo(c echo.Context) error {
	return c.String(http.StatusOK, "ola")
}
