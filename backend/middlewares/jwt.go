package middleware

import (
	"backend/utils"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tokenString string
		
		// 1. Ambil dari Header
		authHeader := c.Get("Authorization")
		if authHeader != "" && len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			// 2. Fallback ke Cookie
			tokenString = c.Cookies("Authorization")
		}

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization token not found", 
				"status": false,
			})
		}

		// Validasi menggunakan utils milikmu
		token, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token", 
				"status": false,
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Cek Expiration
			expiration := int64(claims["exp"].(float64))
			if time.Now().Unix() > expiration {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Token has expired", 
					"status": false,
				})
			}

			// Simpan ke Context (Fiber menggunakan Locals)
			c.Locals("user", claims["sub"])
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token", 
			"status": false,
		})
	}
}