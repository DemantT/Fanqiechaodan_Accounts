package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteUser(filepath string, user User) error {
	err, userArr := ReadUser(filepath)
	userArr = append(userArr, user)
	b, err := json.Marshal(userArr)
	fmt.Println("b is ", string(b))
	if err != nil {
		fmt.Println("Marshal error:", err)
		return err
	}
	//生成json文件
	fl, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open file err is ", err)
		return err
	}
	defer fl.Close()
	writer := bufio.NewWriter(fl)
	_, err = writer.Write(b)
	fmt.Println("err is ", err)
	// 因为 writer 是带缓存，因此调用 WriterString 方法时，其实
	// 内容是先写入到缓存中，所以需要调用 Flush 方法，将缓冲的数据
	// 真正写入到文件中，否则文件中会没有数据 ！！！
	writer.Flush()
	return nil
}

func ReadUser(filepath string) (error, []User) {
	data, err := ioutil.ReadFile(filepath)
	var v []User
	if err != nil {
		return err, nil
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("unmarsh err is ", err)
		return err, nil
	}
	fmt.Println("v is ", v)
	return nil, v
}
