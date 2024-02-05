package helper

import (
	"errors"
	"testtry2/models"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func ValidateList(user models.User) (errs []string) {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(user)
	errorList := TranslateErrors(err, trans)
	for _, e := range errorList {
		errs = append(errs, e.Error())
	}
	return
}

func TranslateErrors(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := errors.New(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
