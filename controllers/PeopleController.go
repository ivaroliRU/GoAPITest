package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GET api/people/
func GetPeople(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// POST api/people/
func AddPerson(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// GET api/people/:id
func GetPersonByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// DELETE api/people/:id
func DeletePersonByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// PUT api/people/:id
func UpdatePersonByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}