package middlewares

import (
	"github.com/keanutaufan/anitrackr-server/platform/firebase_app"
	"github.com/labstack/echo/v4"
	"strings"
)

func FirebaseAuthMiddleware(fc *firebase_app.FirebaseClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authHeader := ctx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return nil
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				return nil
			}
			idToken := parts[1]

			token, err := fc.Auth.VerifyIDToken(ctx.Request().Context(), idToken)
			if err != nil {
				return err
			}

			ctx.Set("firebase_uid", token.UID)
			return next(ctx)
		}
	}
}
