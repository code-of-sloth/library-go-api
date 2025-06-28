package models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateBookReq struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Desc   string `json:"desc"`
	Sku    string `json:"sku"`
}

func (data CreateBookReq) Validate() (err error) {
	err = validation.ValidateStruct(&data,
		validation.Field(&data.Author, validation.Required, validation.Length(3, 25), is.Alpha),
		validation.Field(&data.Genre, validation.Required, validation.Length(3, 20)),
		validation.Field(&data.Name, validation.Required, validation.Length(1, 25)),
		validation.Field(&data.Desc, validation.Required, validation.Length(1, 250)),
		validation.Field(&data.Sku, validation.Required, validation.Length(15, 15), is.Alphanumeric),
	)

	if err != nil {
		err = fmt.Errorf("createbookreq:error during validation %w", err)
		return
	}

	return
}

type RentBookReq struct {
	Sku    string `json:"sku"`
	UserID string `json:"userID"`
}

func (data RentBookReq) Validate() (err error) {
	err = validation.ValidateStruct(&data,
		validation.Field(&data.UserID, validation.Required, validation.Length(15, 15)),
		validation.Field(&data.Sku, validation.Required, validation.Length(15, 15), is.Alphanumeric),
	)

	if err != nil {
		err = fmt.Errorf("rentbookreq:error during validation %w", err)
		return
	}

	return
}

type FetchBookReq struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNum"`
}

func (data FetchBookReq) Validate() (err error) {
	err = validation.ValidateStruct(&data,
		validation.Field(&data.PageNum, validation.Required, validation.Min(1)),
		validation.Field(&data.PageSize, validation.Required, validation.In(10, 20, 50, 100)),
	)

	if err != nil {
		err = fmt.Errorf("fetchbookreq:error during validation %w", err)
		return
	}

	return
}

type ReturnBookReq struct {
	UserID string `json:"userID"`
	BookID string `json:"bookID"`
}

func (data ReturnBookReq) Validate() (err error) {
	err = validation.ValidateStruct(&data,
		validation.Field(&data.UserID, validation.Required, validation.Length(15, 15)),
		validation.Field(&data.BookID, validation.Required, validation.Length(15, 15)),
	)

	if err != nil {
		err = fmt.Errorf("returnbookreq:error during validation %w", err)
		return
	}

	return
}
