package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wyw14/cry038/internal/application"
	"github.com/wyw14/cry038/internal/middleware"
	"go.uber.org/zap"
	"net/http"
)

type lockInput struct {
	Actor string `json:"actor" validate:"required"`
}

func Routes(app *application.Preparation, logger *zap.Logger) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Trace())
	r.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "healthy"}) })
	r.GET("/readyz", func(c *gin.Context) { c.Status(204) })
	v := validator.New()
	r.POST("/api/v1/sessions/:id/lock", func(c *gin.Context) {
		var in lockInput
		if c.ShouldBindJSON(&in) != nil || v.Struct(in) != nil {
			c.JSON(422, gin.H{"code": "INVALID_ACTOR", "message": "需要确认人", "request_id": c.GetString("request_id")})
			return
		}
		s, err := app.Lock(c, c.Param("id"), in.Actor)
		if err != nil {
			logger.Info("lock denied", zap.Error(err))
			c.JSON(200, gin.H{"status": "locked_with_warning", "message": err.Error(), "request_id": c.GetString("request_id")})
			return
		}
		c.JSON(200, s)
	})
	return r
}
