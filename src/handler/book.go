package handler

import (
	"LibraryGo/src/book"
	"LibraryGo/src/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddBook(c echo.Context) (err error) {
	req := new(models.CreateBookReq)
	err = c.Bind(req)
	if err != nil {
		err = fmt.Errorf("createbookreq:error binding %w", err)
		return httpErrResp(err, c)
	}

	err = c.Validate(req)
	if err != nil {
		err = fmt.Errorf("createbookreq:error validating %w", err)
		return httpErrResp(err, c)
	}

	resp, err := book.AddNewBook(req.Name, req.Author, req.Genre, req.Desc, req.Sku)
	if err != nil {
		err = fmt.Errorf("createbookreq:error adding new book %w", err)
		return httpErrResp(err, c)
	}

	return c.JSON(http.StatusOK, resp)
}

func RemoveBook(c echo.Context) (err error) {
	bookID := c.QueryParam("bookid")
	err = book.RemoveBookByID(bookID)
	if err != nil {
		err = fmt.Errorf("removebook:error deleting book with id %s: %w", bookID, err)
		return httpErrResp(err, c)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Deleted Successfully"})
}

func FetchBooks(c echo.Context) (err error) {
	req := new(models.FetchBookReq)
	err = c.Bind(req)
	if err != nil {
		err = fmt.Errorf("fetchbook:error binding %w", err)
		return httpErrResp(err, c)
	}

	err = c.Validate(req)
	if err != nil {
		err = fmt.Errorf("fetchbook:error validating %w", err)
		return httpErrResp(err, c)
	}

	resp, err := book.GetAvailableBooks(req.PageNum, req.PageSize)
	if err != nil {
		err = fmt.Errorf("fetchbook:error fetching book with pagesize:%d pagenum:%d  %w", req.PageSize, req.PageNum, err)
		return httpErrResp(err, c)
	}

	return c.JSON(http.StatusOK, resp)
}

func RentBook(c echo.Context) (err error) {
	req := new(models.RentBookReq)
	err = c.Bind(req)
	if err != nil {
		err = fmt.Errorf("rentbookreq:error binding %w", err)
		return httpErrResp(err, c)
	}

	err = c.Validate(req)
	if err != nil {
		err = fmt.Errorf("rentbookreq:error validating %w", err)
		return httpErrResp(err, c)
	}

	err = book.LendBookToUser(req.UserID, req.Sku)
	if err != nil {
		err = fmt.Errorf("rentbookreq:error renting book %w", err)
		return httpErrResp(err, c)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Updated Successfully"})
}

func ReturnBook(c echo.Context) (err error) {
	req := new(models.ReturnBookReq)
	err = c.Bind(req)
	if err != nil {
		err = fmt.Errorf("returnbookreq:error binding %w", err)
		return httpErrResp(err, c)
	}

	err = c.Validate(req)
	if err != nil {
		err = fmt.Errorf("returnbookreq:error validating %w", err)
		return httpErrResp(err, c)
	}

	err = book.RetunBookByUser(req.UserID, req.BookID)
	if err != nil {
		err = fmt.Errorf("returnbookreq:error returning book %w", err)
		return httpErrResp(err, c)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Updated Successfully"})
}
