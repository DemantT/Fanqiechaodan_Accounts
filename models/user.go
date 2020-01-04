package models

import (
	"errors"
	"fmt"
)

var (
	UserList map[string]*User
)

// 所有的返回值都用User结构体 TODO

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", "mytoken"}
	UserList["user_11111"] = &u
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"token"`
}

//type Profile struct {
//	Gender  string
//	Age     int
//	Address string
//	Email   string
//}

func AddUser(u User) (*User, error) {
	u.Id = u.Username
	//u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	// 此处调用融云接口 获取用户token 然后将id和token返回给前端 TODO
	rongyunUser, err := CreateUser(u.Id)
	if err != nil {
		fmt.Println("create user err is ", err)
		return nil, err
	}
	u.Profile = rongyunUser.Token
	u.Id = rongyunUser.UserID
	fmt.Println("user is ", u)
	UserList[u.Id] = &u
	return &u, nil
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

//func UpdateUser(uid string, uu *User) (a *User, err error) {
//	if u, ok := UserList[uid]; ok {
//		if uu.Username != "" {
//			u.Username = uu.Username
//		}
//		if uu.Password != "" {
//			u.Password = uu.Password
//		}
//		if uu.Profile.Age != 0 {
//			u.Profile.Age = uu.Profile.Age
//		}
//		if uu.Profile.Address != "" {
//			u.Profile.Address = uu.Profile.Address
//		}
//		if uu.Profile.Gender != "" {
//			u.Profile.Gender = uu.Profile.Gender
//		}
//		if uu.Profile.Email != "" {
//			u.Profile.Email = uu.Profile.Email
//		}
//		return u, nil
//	}
//	return nil, errors.New("User Not Exist")
//}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
