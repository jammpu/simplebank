package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/jammpu/simplebank/util"
)

var validCurrency validator.Func = func(fieldLever validator.FieldLevel) bool {
	if currency, ok := fieldLever.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
