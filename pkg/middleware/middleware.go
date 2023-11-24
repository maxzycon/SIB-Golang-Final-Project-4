package middleware

import (
	"slices"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/config"
	UserService "github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/user/service/impl"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/errors"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/httputil"
)

type GlobalMiddleware struct {
	UserService *UserService.UserService
	Conf        *config.Config
}

func (m *GlobalMiddleware) Protected(roleAccess []uint) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(m.Conf.JWT_SECRET_KEY)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return httputil.WriteErrorResponse(c, errors.ErrUnauthorized)
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			// ---- JWT
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			id, ok := claims["id"].(float64)
			if !ok {
				return httputil.WriteErrorResponse(ctx, errors.ErrUnauthorized)
			}

			// --- get user by id
			userRow, err := m.UserService.GetUserByIdToken(ctx.Context(), uint(id))
			if err != nil {
				log.Errorf("err getUserByIdToken %d", uint(id))
				return httputil.WriteErrorResponse(ctx, errors.ErrUnauthorized)
			}

			if !slices.Contains(roleAccess, userRow.Role) {
				return httputil.WriteErrorResponse(ctx, errors.ErrForbidden)
			}

			resp := &authutil.UserClaims{
				ID:       userRow.ID,
				FullName: userRow.FullName,
				Email:    userRow.Email,
				Role:     userRow.Role,
				Balance:  userRow.Balance,
			}
			// ----- set user to all request with protected
			ctx.Locals("user_profile", resp)
			return ctx.Next()
		},
	})
}
