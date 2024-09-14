package http

import (
	"billing_service/billing/wallet/model"
	"billing_service/internal/http_errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Delivery struct {
	uc model.UseCase
}

func NewDelivery(uc model.UseCase) *Delivery {
	return &Delivery{uc: uc}
}

func (d *Delivery) UpdateBalance(ctx echo.Context) (err error) {
	var (
		req model.UpdateHttpRequest
		res *model.UpdateHttpResponse
	)

	defer func() {
		if err != nil {
			if sendError := ctx.JSON(http_errors.GetCodeAndResponse(err)); sendError != nil {
				err = fmt.Errorf("sending error response: %w: %w", sendError, err)
			}

			return
		}

		res.Status = http.StatusOK

		if sendResponse := ctx.JSON(http.StatusOK, res); sendResponse != nil {
			err = fmt.Errorf("sending response: %w: %w", sendResponse, err)
		}
	}()

	if err = ctx.Bind(&req); err != nil {
		return fmt.Errorf("Delivery.UpdateBalance #1: %w", http_errors.ErrBadRequest)
	}

	if req.Sum < 0 {
		return fmt.Errorf("Delivery.UpdateBalance #2: %w", http_errors.ErrBadRequest)
	}

	res, err = d.uc.UpdateBalance(req)
	if err != nil {
		return fmt.Errorf("Delivery.UpdateBalance #3: %w", err)
	}

	return nil
}

func (d *Delivery) GetBalance(ctx echo.Context) (err error) {
	var (
		req model.GetHttpRequest
		res *model.GetHttpResponse
	)

	defer func() {
		if err != nil {
			if sendError := ctx.JSON(http_errors.GetCodeAndResponse(err)); sendError != nil {
				err = fmt.Errorf("sending error response: %w: %w", sendError, err)
			}

			return
		}

		res.Status = http.StatusOK

		if sendResponse := ctx.JSON(http.StatusOK, res); sendResponse != nil {
			err = fmt.Errorf("sending response: %w: %w", sendResponse, err)
		}
	}()

	if err = ctx.Bind(&req); err != nil {
		return fmt.Errorf("Delivery.GetBalance #1: %w", err)
	}

	res, err = d.uc.GetBalance(req)
	if err != nil {
		return fmt.Errorf("Delivery.GetBalance #2: %w", err)
	}

	return nil
}
