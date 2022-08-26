package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	//第一种写入方式1.1
	file, err := os.OpenFile("../file1.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666) // 创建，写入,清空
	// file, err := os.OpenFile("../file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //追加
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 9; i++ {
		file.WriteString("写入" + strconv.Itoa(i) + "\r\n") //换行\r\n
	}
	//方式1.2
	var str = "1213131 byte"
	file.Write([]byte(str))
	//第二种bufio 写入
	writer := bufio.NewWriter(file)
	writer.WriteString("bufio 写入文件")
	writer.Flush()
	//第三种写入方式
	var str2 = "第三种方式"
	err2 := ioutil.WriteFile("../file3.txt", []byte(str2), 0666)
	if err != nil {
		fmt.Println(err2)
		return
	}
}
