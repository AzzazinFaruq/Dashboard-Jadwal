package middleware

// import (
// 	"backend/utils"
// 	"github.com/gofiber/fiber/v2"
// )

// func APIRateLimit(c *fiber.Ctx) error {
// 	ip := c.IP() // Fiber IP helper
// 	if !apiLimiter.Allow(ip) {
// 		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
// 			"error":   "Rate limit exceeded",
// 			"message": "Too many requests. Please try again later.",
// 			"status":  false,
// 		})
// 	}
// 	return c.Next()
// }

// func LoginRateLimit(c *fiber.Ctx) error {
// 	ip := c.IP()
// 	if !loginLimiter.Allow(ip) {
// 		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
// 			"error":   "Rate limit exceeded",
// 			"message": "Too many login attempts. Please try again later.",
// 			"status":  false,
// 		})
// 	}
// 	return c.Next()
// }