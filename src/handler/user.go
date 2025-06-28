package handler

import (
	"LibraryGo/src/models"
	"LibraryGo/src/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) (err error) {
	req := new(models.CreateUserReq)
	err = c.Bind(req)
	if err != nil {
		err = fmt.Errorf("createuser:error binding %w", err)
		return httpErrResp(err, c)
	}

	err = c.Validate(req)
	if err != nil {
		err = fmt.Errorf("createuser:error validating %w", err)
		return httpErrResp(err, c)
	}

	resp, err := user.AddNewUser(req.Name, req.MobileNo)
	if err != nil {
		err = fmt.Errorf("createuser:error adding new user %w", err)
		return httpErrResp(err, c)
	}

	return c.JSON(http.StatusOK, resp)
}

func FetchUser(c echo.Context) (err error) {
	mobileNum := c.QueryParam("mobilenum")
	resp, err := user.FetchUserByMobileNum(mobileNum)
	if err != nil {
		err = fmt.Errorf("fetchuser:error adding new user %w", err)
		return httpErrResp(err, c)
	}

	return c.JSON(http.StatusOK, resp)
}

func UpdateUser(c echo.Context) (err error) {
	req := new(models.UpdateUserReq)
	err = c.Bind(req)
	if err != nil {
		err = fmt.Errorf("updateuser:error binding %w", err)
		return httpErrResp(err, c)
	}

	err = c.Validate(req)
	if err != nil {
		err = fmt.Errorf("updateuser:error validating %w", err)
		return httpErrResp(err, c)
	}

	var data user.UserInfo
	data.MobileNo = req.MobileNo
	data.Name = req.Name
	data.UserID = req.UserID
	err = user.UpdateUserInfo(data)
	if err != nil {
		err = fmt.Errorf("updateuser:error updating user %w", err)
		return httpErrResp(err, c)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Updated Successfully"})
}

func DeleteUser(c echo.Context) (err error) {
	mobileNum := c.QueryParam("mobilenum")
	err = user.DeleteUserByMobileNum(mobileNum)
	if err != nil {
		err = fmt.Errorf("deleteuser:error adding new user %w", err)
		return httpErrResp(err, c)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Deleted Successfully"})
}
