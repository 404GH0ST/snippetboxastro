package main

import "github.com/labstack/echo/v4"

func SendError(c echo.Context, status int, message string) error {
	return c.JSON(status, ErrorResponse{Message: message})
}
