package validators

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField	string
	Tag			string
}

var validate = validator.New()

func ValidateStruct (postBody interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(postBody)
	log.Println("pkg validators func ValidateStruct err ",err)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			log.Println ("pkg validators func ValidateStruct element ",element.FailedField, element.Tag)
			errors = append(errors, &element)
		}
	}
	return errors
}

type AuthorAddPostBody {
	Title string `json:"title" validate:"required"`
}