package handler

import (
	"net/http"
	"grey/internal/domain"
	"grey/internal/handler/response"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) productRoutes(api *gin.RouterGroup) {
	product := api.Group("/product")
	{
		product.POST("/create", h.createProduct)
		product.GET("/", h.listProduct)
		product.GET("/:product_id", h.getProduct)
		product.POST("/:product_id/update-price", h.updatePrice)
	}
}

// @Summary Create product
// @Description Create product
// @Accept  json
// @Produce  json
// @Param input body domain.ProductInput true "ProductInput"
// @Router /api/product/create [post]
func (h *Handler) createProduct(ctx *gin.Context) {
	var input domain.ProductInput

	if err := ctx.BindJSON(&input); err != nil {
		log.Error().Err(err).Msg("")
		response.NewErrorResponse(ctx, http.StatusBadRequest, MsgBadRequest)

		return
	}

	err := h.s.Product.CreateProduct(ctx.Request.Context(), input)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, MsgInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, response.MsgWrapper{Message: "product created"})
}

// @Summary Show list of products
// @Description Show list of products
// @Accept  json
// @Produce  json
// @Router /api/product [get]
func (h *Handler) listProduct(ctx *gin.Context) {
	res, err := h.s.Product.ListProduct(ctx.Request.Context())
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, MsgInternalServerError)

		return
	}
	
	ctx.JSON(http.StatusOK, res)
}

// @Summary Get product
// @Description Get product
// @Accept  json
// @Produce  json
// @Param product_id path int true "int"
// @Router /api/product/{product_id} [get]
func (h *Handler) getProduct(ctx *gin.Context) {
	productId, err := RequiredIntParam(ctx, "product_id")
	if err != nil {
		log.Error().Err(err).Msg("")
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	res, err := h.s.Product.GetProduct(ctx.Request.Context(), productId)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, MsgInternalServerError)

		return
	}
	
	ctx.JSON(http.StatusOK, res)
}

// @Summary Update product price
// @Description Update product price
// @Accept  json
// @Produce  json
// @Param input body domain.ProductPriceInput true "ProductPriceInput"
// @Param product_id path int true "int"
// @Router /api/product/{product_id}/update-price [post]
func (h *Handler) updatePrice(ctx *gin.Context) {
	var input domain.ProductPriceInput

	if err := ctx.BindJSON(&input); err != nil {
		log.Error().Err(err).Msg("")
		response.NewErrorResponse(ctx, http.StatusBadRequest, MsgBadRequest)

		return
	}

	productId, err := RequiredIntParam(ctx, "product_id")
	if err != nil {
		log.Error().Err(err).Msg("")
		response.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	err = h.s.Product.UpdatePrice(ctx.Request.Context(), productId, input)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, MsgInternalServerError)

		return
	}
	
	ctx.JSON(http.StatusOK, response.MsgWrapper{Message: "product price updated"})
}
