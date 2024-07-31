package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/404GH0ST/snippetboxastro/internal/models"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreatedResponse struct {
	ID int `json:"id"`
}

type ViewRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Expires string `json:"expires"`
}

type ViewResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Created string `json:"created"`
	Expires string `json:"expires"`
}

func (app *Application) snippetView(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return SendError(c, http.StatusBadRequest, "Invalid ID")
	}

	snippet, err := app.model.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			return SendError(c, http.StatusNotFound, "No Matching Record Found")
		} else {
			c.Logger().Error(err)
			return SendError(c, http.StatusInternalServerError, "Internal Server Error")
		}
	}

	response := &ViewResponse{
		snippet.ID,
		snippet.Title,
		snippet.Content,
		snippet.Created.String(),
		snippet.Expires.String(),
	}

	return c.JSON(http.StatusOK, response)
}

func (app *Application) snippetCreate(c echo.Context) error {
	var request ViewRequest
	if err := c.Bind(&request); err != nil {
		return SendError(c, http.StatusBadRequest, "Invalid Request Body")
	}
	id, err := app.model.Insert(request.Title, request.Content, request.Expires)
	if err != nil {
		c.Logger().Error(err)
		return SendError(c, http.StatusInternalServerError, "Internal Server Error")
	}

	response := CreatedResponse{
		ID: id,
	}

	return c.JSON(http.StatusOK, response)
}

func (app *Application) snippetLatest(c echo.Context) error {
	snippets, err := app.model.Latest()
	if err != nil {
		c.Logger().Error(err)
		return SendError(c, http.StatusInternalServerError, "Internal Server Error")
	}

	var response []ViewResponse
	for _, snippet := range snippets {
		response = append(response, ViewResponse{
			ID:      snippet.ID,
			Title:   snippet.Title,
			Content: snippet.Content,
			Created: snippet.Created.String(),
			Expires: snippet.Expires.String(),
		})
	}

	return c.JSON(http.StatusOK, response)
}
