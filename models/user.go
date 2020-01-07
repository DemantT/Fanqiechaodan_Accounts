package models

import (
	"errors"
	"fmt"

	"gopkg.in/gomail.v2"
)

var (
	UserList map[string]*User
)

func init() {
	err, userArr := ReadUser("/Users/liuchuan/myowncode/testcode/test.json")
	if err != nil {
		fmt.Println("read user err is ", err)
	}
	UserList = make(map[string]*User)
	for _, user := range userArr {
		UserList[user.Id] = &user
	}
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"token"`
}

type Mail struct {
	To    string `json:"to"`
	Title string `json:"title"`
	Msg   string `json:"message"`
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

func Login(username, password string) (*User, error, bool) {
	fmt.Println("user list is ", UserList)
	u := new(User)
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			fmt.Println("get user")
			return u, nil, true
		}
	}
	user, err := CreateUser(username)
	if err != nil {
		fmt.Println("create user err is ", err)
		return nil, err, false
	}
	u.Password = password
	u.Username = username
	u.Profile = user.Token
	u.Id = user.UserID
	fmt.Println("user is ", u)
	UserList[u.Id] = u
	return u, err, false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

func SendMail(userId string, mail Mail) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "18291169093@163.com", "me") // 发件人
	m.SetHeader("To",                                       // 收件人
		m.FormatAddress(mail.To, "test"),
	)
	m.SetHeader("Subject", mail.Title) // 主题
	m.SetBody("text/html", mail.Msg)   // 正文

	d := gomail.NewPlainDialer("smtp.163.com", 465, "18291169093@163.com", "lc931027") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	filePath := "/Users/liuchuan/myowncode/testcode/" + userId + "+mail.json"
	err := WriteJson(filePath, mail)
	if err != nil {
		fmt.Println("write json err is ", err)
		return err
	}
	return nil
}
