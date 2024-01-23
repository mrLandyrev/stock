package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrLandyrev/stock/internal/errors"
	"github.com/mrLandyrev/stock/internal/models"
	"github.com/mrLandyrev/stock/internal/uescases/get_balance"
	"github.com/mrLandyrev/stock/internal/uescases/reserve"
	"github.com/mrLandyrev/stock/internal/uescases/unreserve"
)

func BuildHandlers(
	getBalanceUseCase *get_balance.GetBalanceUseCase,
	reserveUseCase *reserve.ReserveUseCase,
	unreserveUseCase *unreserve.UnreserveUseCase,
) http.Handler {
	r := gin.Default()
	r.GET("/balance/:stockId", func(c *gin.Context) {
		result, err := getBalanceUseCase.Execute(models.StockId(c.Param("stockId")))

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, result)
	})
	r.POST("/reserve", func(c *gin.Context) {
		var productIds []string
		err := c.BindJSON(&productIds)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		fmt.Println(productIds)
		reserveProducts := make(models.ReserveProducts, 0)
		for _, productId := range productIds {
			reserveProducts[models.ProductId(productId)]++
		}
		err = reserveUseCase.Execute(reserveProducts)

		switch err {
		case nil:
			c.Status(http.StatusOK)
		case errors.ErrProductOutOfStock:
			c.Status(http.StatusBadRequest)
		default:
			c.Status(http.StatusInternalServerError)
		}
	})
	r.POST("/unreserve", func(c *gin.Context) {
		var productIds []string
		err := c.BindJSON(&productIds)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		unreserveProducts := make(models.UnreserveProducts, 0)
		for _, productId := range productIds {
			unreserveProducts[models.ProductId(productId)]++
		}
		err = unreserveUseCase.Execute(unreserveProducts)

		switch err {
		case nil:
			c.Status(http.StatusOK)
		case errors.ErrProductReserveNotFound:
			c.Status(http.StatusBadRequest)
		default:
			c.Status(http.StatusInternalServerError)
		}
	})

	return r
}
