package handler

import (
	"grey/internal/handler/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) orderRoutes(api *gin.RouterGroup) {
	order := api.Group("/order")
	{
		order.GET("/", h.listOrder)
		order.GET("/:order_id", h.detailOrder)
	}
}

// @Summary Show orders for user
// @Description Show list of orders for user with hard coded userId (handler.go -> const)
// @Accept  json
// @Produce  json
// @Router /api/order [get]
func (h *Handler) listOrder(ctx *gin.Context) {
	res, err := h.s.Order.ListOrder(ctx.Request.Context(), userId)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, MsgInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary Show order detail for user
// @Description Show order details for user with hard coded userId(handler.go -> const)
// @Accept  json
// @Produce  json
// @Param order_id path int true "int"
// @Router /api/order/{order_id} [get]
func (h *Handler) detailOrder(ctx *gin.Context) {
	orderId, err := RequiredIntParam(ctx, "order_id")
	if err != nil {
		log.Error().Err(err)
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	res, err := h.s.Order.DetailOrder(ctx.Request.Context(), userId, orderId)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, MsgInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, res)
}
