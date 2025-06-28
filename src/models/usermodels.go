package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateUserReq struct {
	Name     string `json:"name"`
	MobileNo string `json:"mobileNo"`
}

func (data CreateUserReq) Validate() (err error) {
	err = validation.ValidateStruct(&data,
		validation.Field(&data.Name, validation.Required, validation.Length(3, 25), is.Alpha),
		validation.Field(&data.MobileNo, validation.Required, validation.Length(10, 10), is.Digit),
	)

	if err != nil {
		err = fmt.Errorf("createuserreq:error during validation %w", err)
		return
	}

	return
}

type UpdateUserReq struct {
	Name     string `json:"name"`
	MobileNo string `json:"mobileNo"`
	UserID   string `json:"userID"`
}

func (data UpdateUserReq) Validate() (err error) {
	err = validation.ValidateStruct(&data,
		validation.Field(&data.UserID, validation.Required, validation.Length(15, 15)),
		validation.Field(&data.Name, validation.Required, validation.Length(3, 25), is.Alpha),
		validation.Field(&data.MobileNo, validation.Required, validation.Length(10, 10), is.Digit),
	)

	if err != nil {
		err = fmt.Errorf("updateuserreq:error during validation %w", err)
		return
	}

	return
}
