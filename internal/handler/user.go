package handler

import (
	"errors"
	"grey/internal/domain"
	"grey/internal/handler/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) userRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
	}
}

// @Summary Sign up
// @Description Sign up
// @Accept  json
// @Produce  json
// @Param input body domain.UserSignUpInput true "UserSignUpInput"
// @Router /api/auth/sign-up [post]
func (h *Handler) signUp(ctx *gin.Context) {
	var input domain.UserSignUpInput

	if err := ctx.BindJSON(&input); err != nil {
		log.Error().Err(err).Msg("")
		response.NewErrorResponse(ctx, http.StatusBadRequest, MsgBadRequest)

		return
	}

	err := h.s.User.SignUp(ctx.Request.Context(), input)
	if err != nil {
		if errors.Is(err, domain.ErrRecordAlreadyExists) {
			response.NewErrorResponse(ctx, http.StatusBadRequest, "user already exists")

			return
		}

		response.NewErrorResponse(ctx, http.StatusInternalServerError, MsgInternalServerError)

		return
	}

	ctx.JSON(http.StatusCreated, response.MsgWrapper{Message: "account created"})
}