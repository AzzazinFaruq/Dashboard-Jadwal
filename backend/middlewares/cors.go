package middleware

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSMiddleware() fiber.Handler {
	feURL := os.Getenv("FE_URL")
	if feURL == "" {
		feURL = "http://localhost:3000"
	}

	// Inisialisasi CORS bawaan Fiber
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     feURL,
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	})

	return func(c *fiber.Ctx) error {
		// Jalankan CORS bawaan
		if err := corsMiddleware(c); err != nil {
			return err
		}

		// Security Headers (Sesuai kode Gin milikmu)
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "DENY")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Set("Content-Security-Policy", "default-src 'self'")

		// Cache Control untuk Auth (Sesuai kode Gin milikmu)
		path := c.Path()
		if path == "/login" || path == "/register" {
			c.Set("Cache-Control", "no-store, no-cache, must-revalidate, private")
			c.Set("Pragma", "no-cache")
			c.Set("Expires", "0")
		}

		// Handle OPTIONS (Fiber menangani ini secara internal di middleware CORS, 
		// tapi jika ingin manual seperti kodemu:)
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	}
}