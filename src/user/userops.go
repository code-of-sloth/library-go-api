package user

import (
	"LibraryGo/src/config"
	"LibraryGo/src/utils"
	"context"
	"fmt"
)

func AddNewUser(name, mobileNo string) (data UserInfo, err error) {
	data.UserID = utils.GenerateRandomChar("user", 10)
	_, err = config.DbConn.Exec(context.Background(), "INSERT INTO library.users (userid,name,mobile,createdat,updatedat,isactive) VALUES ($1,$2,$3,now(),now(),true)", data.UserID, name, mobileNo)
	if err != nil {
		err = fmt.Errorf("addnewuser:error saving new user %w", err)
		return
	}

	data.Name = name
	data.MobileNo = mobileNo
	return
}

func FetchUserByMobileNum(mobileNum string) (data *UserInfo, err error) {
	if mobileNum == "" {
		err = fmt.Errorf("fetchuserbymobilenum:mobile num cannot be empty")
		return
	}

	rows, err := config.DbConn.Query(context.Background(), "SELECT userid,name FROM library.users WHERE mobile=$1 AND isactive=true LIMIT 1", mobileNum)
	if err != nil {
		err = fmt.Errorf("fetchuserbymobilenum:error querying users for mobile %s : %w", mobileNum, err)
		return
	}

	defer rows.Close()
	if rows.Next() {
		var usr UserInfo
		err = rows.Scan(&usr.UserID, &usr.Name)
		if err != nil {
			err = fmt.Errorf("fetchuserbymobilenum:error reading users for mobile %s : %w", mobileNum, err)
			return
		}

		usr.MobileNo = mobileNum
		data = &usr
	} else {
		err = fmt.Errorf("fetchuserbymobilenum:unable to find any active users for mobile num %s", mobileNum)
		return
	}
	return
}

func UpdateUserInfo(data UserInfo) (err error) {
	_, err = config.DbConn.Exec(context.Background(), "UPDATE library.users SET name=$1,mobile=$2,updatedat=now() WHERE isactive=true AND userid=$3", data.Name, data.MobileNo, data.UserID)
	if err != nil {
		err = fmt.Errorf("updateuserinfo:error updating user info for id %s : %w", data.UserID, err)
		return
	}
	return
}

func DeleteUserByMobileNum(mobileNum string) (err error) {
	_, err = config.DbConn.Exec(context.Background(), "UPDATE library.users SET isactive=false,deletedat=now(),updatedat=now() WHERE isactive=true AND mobile=$1", mobileNum)
	if err != nil {
		err = fmt.Errorf("deleteuser:error deleting user info for mobile %s : %w", mobileNum, err)
		return
	}

	return
}

func FetchUserByID(userID string) (data *UserInfo, err error) {
	if userID == "" {
		err = fmt.Errorf("fetchuserbyid:userid cannot be empty")
		return
	}

	rows, err := config.DbConn.Query(context.Background(), "SELECT mobile,name FROM library.users WHERE userid=$1 AND isactive=true LIMIT 1", userID)
	if err != nil {
		err = fmt.Errorf("fetchuserbyid:error querying users by userid %s : %w", userID, err)
		return
	}

	defer rows.Close()
	if rows.Next() {
		var usr UserInfo
		err = rows.Scan(&usr.MobileNo, &usr.Name)
		if err != nil {
			err = fmt.Errorf("fetchuserbyid:error reading users for userid %s : %w", userID, err)
			return
		}

		usr.UserID = userID
		data = &usr
	} else {
		err = fmt.Errorf("fetchuserbyid:unable to find any active users for userid %s", userID)
		return
	}
	return
}
