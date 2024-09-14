package wallet

import (
	"billing_service/billing/wallet/http"
	"billing_service/billing/wallet/repository"
	"billing_service/billing/wallet/usecase"
	"github.com/labstack/echo/v4"
)

func InitRouting(group *echo.Group) {
	repo := repository.NewRepository()
	uc := usecase.NewUseCase(repo)
	delivery := http.NewDelivery(uc)

	balanceGroup := group.Group("/balance")
	balanceGroup.POST("/update", delivery.UpdateBalance)
	balanceGroup.GET("/get", delivery.GetBalance)
}
