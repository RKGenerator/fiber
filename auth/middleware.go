package auth

import (
	"os"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
)

func OptionalJWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: os.Getenv("JWT_SECRET"),
		},
		ErrorHandler: func(c fiber.Ctx, err error) error {
			println(err)
			return c.Next()
		},
		SuccessHandler: func(c fiber.Ctx) error {
			token := jwtware.FromContext(c)
			println(token)
			return c.Next()
		},
	})
}
