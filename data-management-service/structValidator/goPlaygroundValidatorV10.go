package validator

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type playgroundValidator struct {
	validate *validator.Validate
}

func NewValidator() Validator {
	log.Println("Creating new validator.....")
	v := validator.New()
	v.RegisterValidation("youtubeTime", validateYoutubeTime)
	log.Println("Success creating new validator!")
	return &playgroundValidator{
		validate: v,
	}
}

func validateYoutubeTime(fl validator.FieldLevel) bool {
	_, err := time.Parse("15:04:05", fl.Field().String())
	return err == nil
}

func (v *playgroundValidator) GetValidator() *validator.Validate {
	return v.validate
}

func (v *playgroundValidator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}
