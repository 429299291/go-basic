package calcs

import "fmt"

func init() { //	init 在导入包时自动执行，没有参数也没有返回值，最后被导入的包会最先执行
	fmt.Println("calc init...")
}
func Add(x, y int) int { //首字母大写公开
	return x + y
}
func Sub(x, y int) int {
	return x - y
}
