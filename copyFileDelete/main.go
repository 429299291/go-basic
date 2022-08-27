package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func copyFile(srcFileName string, dstFileName string) (err error) {
	byteStr, err := ioutil.ReadFile(srcFileName)
	if err != nil {
		return err
	}
	err2 := ioutil.WriteFile(dstFileName, byteStr, 0666)
	if err2 != nil {
		return err2
	}
	return nil
}
func copyFile2(srcFileName string, dstFileName string) (err error) {
	Sfile, err := os.Open(srcFileName)
	Dfile, err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	defer Sfile.Close()
	defer Dfile.Close()
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	var temSlice = make([]byte, 128)
	for {
		n1, err := Sfile.Read(temSlice)
		_, err2 := Dfile.Write(temSlice[:n1])
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err2 != nil {
			return err2
		}
	}
	return nil
}
func main() {
	//复制
	// err := copyFile("../writeFile/file1.txt", "./file1.txt")
	err := copyFile2("../writeFile/file1.txt", "./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建文件夹
	err2 := os.Mkdir("./dir", 0666)
	// err2 := os.MkdirAll("./dir1/dir2", 0666)
	if err != nil {
		fmt.Println(err2)
	}
	//删除文件
	err3 := os.Remove("./file.txt")
	if err != nil {
		fmt.Println(err3)
	}
	//删除文件夹	 os.RemoveAll 删除多个文件
	err4 := os.Remove("./dir")
	if err4 != nil {
		fmt.Println(err3)
	}
	//重命名
	err5 := os.Rename("./dir", "./dirRename")
	if err5 != nil {
		fmt.Println(err3)
	}
}
