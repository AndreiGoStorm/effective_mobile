package middlewares

import "github.com/gofiber/fiber/v2"

// Cors -
func Cors() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderAccessControlAllowOrigin, "*")
		c.Set(fiber.HeaderAccessControlAllowMethods, "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Set(fiber.HeaderAccessControlAllowHeaders, "authorization, origin, content-type, accept")
		c.Set(fiber.HeaderAllow, "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		c.Set(fiber.HeaderContentType, "multipart/form-data; boundary=something")
		c.Status(fiber.StatusOK)
		return c.Next()
	}
}
