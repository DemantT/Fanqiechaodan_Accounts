package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/rongcloud/server-sdk-go/sdk"
)

var mysdk *sdk.RongCloud

func init() {
	mysdk = sdk.NewRongCloud(beego.AppConfig.String("runyunapikey"), beego.AppConfig.String("runyunapisecret"), sdk.WithNumTimeout(2))
	fmt.Println("mysdk created is ", *mysdk)
}

func CreateUser(userId string) (*sdk.User, error) {
	userInfo, err := mysdk.UserRegister(userId, userId, "https://data.photo-ac.com/data/thumbnails/20/20d9791648cd4e8c19c4334c7d1a9f20_w.jpeg")
	if err != nil {
		fmt.Println("UserRegister error is ", err)
		return nil, err
	}

	return &userInfo, nil
}
