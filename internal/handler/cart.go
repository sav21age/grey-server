package handler

import (
	"net/http"
	"grey/internal/domain"
	"grey/internal/handler/response"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) cartRoutes(api *gin.RouterGroup) {
	cart := api.Group("/cart")
	{
		cart.POST("/add", h.addProductToCart)
		cart.GET("/", h.listCart)
		cart.GET("/checkout", h.checkout)
	}
}

// @Summary Add product to cart for user
// @Description Add product to cart for user with hard coded userId (handler.go -> const)
// @Accept  json
// @Produce  json
// @Param input body domain.CartInput true "CartInput"
// @Router /api/cart/add [post]
func (h *Handler) addProductToCart(ctx *gin.Context) {
	var input domain.CartInput

	if err := ctx.BindJSON(&input); err != nil {
		log.Error().Err(err).Msg("")
		response.NewErrorResponse(ctx, http.StatusBadRequest, MsgBadRequest)

		return
	}

	err := h.s.Cart.AddProduct(ctx.Request.Context(), userId, input)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, MsgInternalServerError)

		return
	}

	ctx.JSON(http.StatusCreated, response.MsgWrapper{Message: "product added to cart"})
}

// @Summary Show cart for user
// @Description Show list of items for user with hard coded userId (handler.go -> const)
// @Accept  json
// @Produce  json
// @Router /api/cart [get]
func (h *Handler) listCart(ctx *gin.Context) {
	res, err := h.s.Cart.ListCart(ctx.Request.Context(), userId)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, MsgInternalServerError)

		return
	}
	
	ctx.JSON(http.StatusOK, res)
}

// @Summary Checkout order for user
// @Description Checkout order for user with hard coded userId (handler.go -> const)
// @Accept  json
// @Produce  json
// @Router /api/cart/checkout [get]
func (h *Handler) checkout(ctx *gin.Context) {
	err := h.s.Cart.CheckoutCart(ctx.Request.Context(), userId)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, MsgInternalServerError)

		return
	}
	
	ctx.JSON(http.StatusCreated, response.MsgWrapper{Message: "cart checkout"})
}