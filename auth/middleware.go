package auth

import (
	"os"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
)

func OptionalJWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(os.Getenv("JWT_SECRET")),
		},
		Claims: &AuthClaims{},
		ErrorHandler: func(c fiber.Ctx, err error) error {
			println(err.Error())
			return c.Next()
		},
		SuccessHandler: func(c fiber.Ctx) error {
			token := jwtware.FromContext(c)
			println(token)
			return c.Next()
		},
	})
}

func AutorizedJWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(os.Getenv("JWT_SECRET")),
		},
		Claims: &AuthClaims{},
	})
}
