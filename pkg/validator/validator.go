package validator

import "github.com/go-playground/validator/v10"

// Validate 是全局验证器实例
var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
