package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/zzooman/zapp-server/utils"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	return utils.IsSupportedCurrency(fl.Field().String())
}

