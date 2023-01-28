package facade

import "errors"

var (
	ErrCardNotFound            = errors.New("card not found")
	ErrProductNotFound         = errors.New("product not found")
	ErrNoMoneyOnBalance        = errors.New("no money on balance")
	ErrNotEnoughMoneyOnBalance = errors.New("not enough money on balance")
)
