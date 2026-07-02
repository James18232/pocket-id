package middleware

import (
	"log/slog"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pocket-id/pocket-id/backend/internal/common"
	"github.com/pocket-id/pocket-id/backend/internal/service"
	"github.com/pocket-id/pocket-id/backend/internal/utils/cookie"
)

type JwtAuthMiddleware struct {
	userService *service.UserService
	jwtService  *service.JwtService
}

func NewJwtAuthMiddleware(jwtService *service.JwtService, userService *service.UserService) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{jwtService: jwtService, userService: userService}
}

func (m *JwtAuthMiddleware) Add(adminRequired bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, isAdmin, authenticationMethod, authenticationTime, permittedClients, err := m.Verify(c, adminRequired)
		if err != nil {
			c.Abort()
			_ = c.Error(err)
			return
		}
		slog.InfoContext(c.Request.Context(), "jwt middleware execution tracking",
			"permittedClients", permittedClients,
			"userID", userID,
		)
		c.Set("userID", userID)
		c.Set("userIsAdmin", isAdmin)
		c.Set("authenticationMethod", authenticationMethod)
		c.Set("authenticationTime", authenticationTime)
		c.Set("permittedClients", permittedClients)
		c.Next()
	}
}

func (m *JwtAuthMiddleware) Verify(c *gin.Context, adminRequired bool) (subject string, isAdmin bool, authenticationMethod string, authenticationTime time.Time, permittedClients string, err error) {
	ctx := c.Request.Context()
	slog.InfoContext(ctx, "[DEBUG-VERIFY] Middleware execution started", "adminRequired", adminRequired)

	accessToken, err := c.Cookie(cookie.AccessTokenCookieName)
	if err != nil {
		slog.InfoContext(ctx, "[DEBUG-VERIFY] Access token cookie not found, attempting Authorization header fallback", "cookieError", err.Error())
		var ok bool
		_, accessToken, ok = strings.Cut(c.GetHeader("Authorization"), " ")
		if !ok || accessToken == "" {
			slog.WarnContext(ctx, "[DEBUG-VERIFY] Failed: No token found in Cookie or Authorization header")
			return "", false, "", time.Time{}, "", &common.NotSignedInError{}
		}
	}

	slog.InfoContext(ctx, "[DEBUG-VERIFY] Token located successfully", "tokenLength", len(accessToken))

	token, err := m.jwtService.VerifyAccessToken(accessToken)
	if err != nil {
		slog.WarnContext(ctx, "[DEBUG-VERIFY] Failed: VerifyAccessToken rejected the token", "error", err.Error())
		return "", false, "", time.Time{}, "", &common.NotSignedInError{}
	}

	permittedClients, err = m.jwtService.GetPermittedClients(token)
	if err != nil {
		slog.WarnContext(ctx, "[DEBUG-VERIFY] Failed: GetPermittedClients returned an error", "error", err.Error())
		return "", false, "", time.Time{}, "", &common.NotSignedInError{}
	}
	slog.InfoContext(ctx, "[DEBUG-VERIFY] Extracted permittedClients claim", "permittedClients", permittedClients)

	authenticationMethod, err = m.jwtService.GetAuthenticationMethod(token)
	if err != nil {
		slog.WarnContext(ctx, "[DEBUG-VERIFY] Failed: GetAuthenticationMethod returned an error", "error", err.Error())
		return "", false, "", time.Time{}, "", &common.NotSignedInError{}
	}
	authenticationTime, _ = token.IssuedAt()

	subject, ok := token.Subject()
	if !ok {
		slog.WarnContext(ctx, "[DEBUG-VERIFY] Failed: Subject missing from token claims")
		_ = c.Error(&common.TokenInvalidError{})
		return "", false, "", time.Time{}, "", &common.TokenInvalidError{}
	}
	slog.InfoContext(ctx, "[DEBUG-VERIFY] Token parsed completely", "subject", subject, "authMethod", authenticationMethod)

	user, err := m.userService.GetUser(c, subject)
	if err != nil {
		slog.ErrorContext(ctx, "[DEBUG-VERIFY] Failed: GetUser database/service error", "error", err.Error(), "subject", subject)
		return "", false, "", time.Time{}, "", &common.NotSignedInError{}
	}

	if user.Disabled {
		slog.WarnContext(ctx, "[DEBUG-VERIFY] Failed: Target user is disabled", "subject", subject)
		return "", false, "", time.Time{}, "", &common.UserDisabledError{}
	}

	if adminRequired && !user.IsAdmin {
		slog.WarnContext(ctx, "[DEBUG-VERIFY] Failed: Admin access required but user is not an administrator", "subject", subject)
		return "", false, "", time.Time{}, "", &common.MissingPermissionError{}
	}

	slog.InfoContext(ctx, "[DEBUG-VERIFY] Success: Verification completed entirely",
		"subject", subject,
		"isAdmin", user.IsAdmin,
		"authenticationMethod", authenticationMethod,
		"authenticationTime", authenticationTime,
		"permittedClients", permittedClients,
	)

	return subject, user.IsAdmin, authenticationMethod, authenticationTime, permittedClients, nil
}
