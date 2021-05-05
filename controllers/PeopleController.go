package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GET api/people/
func index(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// POST api/people/
func addPerson(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// GET api/people/:id
func getPersonByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// DELETE api/people/:id
func deletePersonByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// PUT api/people/:id
func updatePersonByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}