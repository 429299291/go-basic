package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//只读	file1, err := os.Open("/Users/zhudongdong/go/src/main/README.md")
	file, err := os.Open("../README.md")
	defer file.Close() //必须得关闭
	fmt.Println(file, err)
	//1.
	var strSlice []byte
	temSlice := make([]byte, 128)
	for {
		fileNum, err := file.Read(temSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		fmt.Println("--", fileNum, err, string(temSlice))
		strSlice = append(strSlice, temSlice[:fileNum]...) //最后一次不够128字节处理
	}
	fmt.Println("---", string(strSlice), "---")
	//2.bufio
	reader := bufio.NewReader(file)
	var fileString string
	for {
		str, err := reader.ReadString('\n')
		if err != io.EOF {
			fileString += str //这个要注意
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fileString += str
	}
	fmt.Println(fileString)
	//3.
	byteStr, err := ioutil.ReadFile("../README.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))
}
