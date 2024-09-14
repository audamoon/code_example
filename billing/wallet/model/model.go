package model

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

const (
	AddOperation      = "add"
	SubtractOperation = "subtract"
)

type UseCase interface {
	UpdateBalance(UpdateHttpRequest) (*UpdateHttpResponse, error)
	GetBalance(GetHttpRequest) (*GetHttpResponse, error)
}

type Repository interface {
	GetCurrentBalance(tx *sql.Tx, walletID int) (int, error)
	UpdateBalance(tx *sql.Tx, walletID, balance int) error
}

type Delivery interface {
	GetBalance(ctx echo.Context) error
	UpdateBalance(ctx echo.Context) error
}

// Update operation

type UpdateHttpRequest struct {
	WalletID  int    `json:"wallet_id"`
	Operation string `json:"operation"`
	Sum       int    `json:"sum"`
}

type UpdateHttpResponse struct {
	Status  int `json:"status"`
	Balance int `json:"balance"`
}

type UpdateUseCaseResponse struct {
	Balance int `json:"balance"`
}

//Get operation

type GetHttpRequest struct {
	WalletID int `json:"wallet_id"`
	Balance  int `json:"balance"`
}

type GetHttpResponse struct {
	Status  int `json:"status"`
	Balance int `json:"balance"`
}

type GetUsecaseResponse struct {
	Balance int `json:"balance"`
}
