package middlewares

import (
	"github.com/keanutaufan/anitrackr-server/pkg/app_error"
	"github.com/keanutaufan/anitrackr-server/platform/firebase_app"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func FirebaseAuthMiddleware(fc *firebase_app.FirebaseClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return app_error.New(http.StatusUnauthorized, "Missing authorization header")
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				return app_error.New(http.StatusUnauthorized, "Invalid authorization header")
			}
			idToken := parts[1]

			token, err := fc.Auth.VerifyIDToken(c.Request().Context(), idToken)
			if err != nil {
				return app_error.New(http.StatusUnauthorized, "Invalid or expired token")
			}

			userId, ok := token.Claims["app_user_id"]
			if !ok {
				return app_error.New(http.StatusUnauthorized, "Invalid token")
			}

			userIdFloat, ok := userId.(float64)
			if !ok {
				return app_error.New(http.StatusUnauthorized, "Invalid token")
			}

			c.Set("firebase_uid", token.UID)
			c.Set("userId", int64(userIdFloat))
			return next(c)
		}
	}
}
