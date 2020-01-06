package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//type User struct {
//	Id       string `json:"id"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//	Profile  string `json:"token"`
//}

func WriteUser(filepath string, user User) error {
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return err
	}
	//生成json文件
	fl, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open file err is ", err)
		return err
	}
	defer fl.Close()
	writer := bufio.NewWriter(fl)
	writer.Write(b)
	// 因为 writer 是带缓存，因此调用 WriterString 方法时，其实
	// 内容是先写入到缓存中，所以需要调用 Flush 方法，将缓冲的数据
	// 真正写入到文件中，否则文件中会没有数据 ！！！
	writer.Flush()
	return nil
}

func ReadUser(filepath string) (error, *User) {
	data, err := ioutil.ReadFile(filepath)
	var v *User
	if err != nil {
		return err, nil
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return err, nil
	}
	return nil, v
}

func main() {
	use1 := User{"1", "2", "3", "4"}
	use2 := User{"2", "2", "3", "4"}
	use3 := User{"3", "2", "3", "4"}
	WriteUser("./test.json", use1)
	WriteUser("./test.json", use2)
	WriteUser("./test.json", use3)
}
