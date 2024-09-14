package usecase

import (
	"billing_service/billing/wallet/model"
	"billing_service/internal/db"
	"billing_service/internal/http_errors"
	"fmt"
)

type UseCase struct {
	repo model.Repository
}

func NewUseCase(repo model.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) UpdateBalance(req model.UpdateHttpRequest) (res *model.UpdateHttpResponse, err error) {
	tx, err := db.Instance.Begin()
	if err != nil {
		return nil, fmt.Errorf("UseCase.UpdateBalance #1: %w", err)
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = fmt.Errorf("UseCase.UpdateBalance tx rollback error: %w: %w", rollbackErr, err)
			}

			return
		}

		if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("UseCase.UpdateBalance tx commit error: %w: %w", commitErr, err)
		}
	}()

	res = new(model.UpdateHttpResponse)

	res.Balance, err = uc.repo.GetCurrentBalance(tx, req.WalletID)
	if err != nil {
		return nil, fmt.Errorf("UseCase.UpdateBalance #2: %w", err)
	}

	switch req.Operation {
	case model.AddOperation:
		res.Balance += req.Sum
	case model.SubtractOperation:
		if res.Balance-req.Sum < 0 {

			return nil, http_errors.ErrNotEnoughMoney
		}

		res.Balance -= req.Sum
	}

	err = uc.repo.UpdateBalance(tx, req.WalletID, res.Balance)
	if err != nil {
		return nil, fmt.Errorf("UseCase.UpdateBalance #3: %w", err)
	}

	return res, nil
}

func (uc *UseCase) GetBalance(req model.GetHttpRequest) (*model.GetHttpResponse, error) {
	var (
		res model.GetHttpResponse
		err error
	)

	tx, err := db.Instance.Begin()
	if err != nil {
		return nil, fmt.Errorf("UseCase.GetBalance #1: %w", err)
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = fmt.Errorf("UseCase.UpdateBalance tx rollback error: %w: %w", rollbackErr, err)
			}

			return
		}

		if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("UseCase.UpdateBalance tx commit error: %w: %w", commitErr, err)
		}
	}()

	res.Balance, err = uc.repo.GetCurrentBalance(tx, req.WalletID)
	if err != nil {
		return nil, fmt.Errorf("UseCase.GetBalance #2: %w", err)
	}

	return &res, nil
}
