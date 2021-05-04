package utils

import "github.com/gofiber/fiber/v2"

func JsonError(code int) fiber.Map {
	return fiber.Map{
		"code": code,
	}
}
