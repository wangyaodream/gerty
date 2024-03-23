package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const SessionKey = "session_id"

type SessionAuth struct {
}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sessionID := ctx.GetHeader(SessionKey)
	// TODO : implement auth
	if sessionID == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "session id is null")
	}
	ctx.Next()
}
