package handler

import (
	"errors"
	"fmt"
	"grey/config"
	"grey/internal/service"
	"strconv"
	"time"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "grey/docs"

	"github.com/gin-gonic/gin"
)

var (
	MsgQueryParamInvalid   = "invalid query parameter"
	MsgInternalServerError = "internal server error"
	MsgBadRequest          = "invalid input body"
)

const (
	userId = 1
)

type Handler struct {
	s   *service.Service
	cfg *config.Config
}

func NewHandler(service *service.Service, config *config.Config) *Handler {
	return &Handler{
		s:   service,
		cfg: config,
	}
}

func (h *Handler) InitRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(gin.LoggerWithFormatter(
		func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))

	// router.Use(
	// 	middleware.Cors(h.cfg),
	// )


	api := router.Group("/api")
	{
		h.userRoutes(api)
		h.productRoutes(api)
		h.cartRoutes(api)
		h.orderRoutes(api)
	}
	return router
}

func RequiredIntParam(ctx *gin.Context, name string) (id int, err error) {
	id, err = strconv.Atoi(ctx.Param(name))
	if err != nil {
		msg := fmt.Sprintf("%s: %s", MsgQueryParamInvalid, name)
		return 0, errors.New(msg)
	}

	return
}