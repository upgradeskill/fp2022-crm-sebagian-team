package domain

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrConflict            = errors.New("Your Item already exist")
	ErrBadParamInput       = errors.New("Given Param is not valid")
	ErrInvalidCredentials  = errors.New("Invalid credentials")
	ErrDuplicateNIP        = errors.New("NIP already exists")
	ErrUserIsNotActive     = errors.New("Your account is not active")
	ErrForbidden           = errors.New("You are not authorized to access this resource")
)

type ErrResponse struct {
	Message string `json:"message"`
}

func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{Message: err.Error()}
}
