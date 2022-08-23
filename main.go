package main

import (
	"encoding/json"
	"fmt"
	"main/calcs" //自定义包，go mod init 项目名
	"sort"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// 函数<<<<<<<<<<<<<<
func func1(x ...int) (int, int) {
	fmt.Print(x)
	return 8, 9
}
func func2(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

type calcType func(int, int) (int, int) //type 自定义类型main.xx
type myFloat = float32                  //类型别名

func calc(x, y int, cd calcType) (int, int) {
	return cd(x, y)
}

// 闭包<<<<<<<<<<<<<<定义在函数内部的函数，连接函数内外,常驻内存，不污染全局
// defer<<<<<<<<<<<<<<延迟处理,逆序执行,return ????????????
//panic 异常，结束执行程序，recover必须在 defer 内,接住异常不终止
func f4(x, y int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover", err)
		}
	}()
	return x / y
}
func f3() (a int) {
	defer func() {
		a++
	}()
	return
}

//指针
func pointFun(x *int) {
	*x = 66
}

//结构体<<<<<<<<<<<<<<<<<<<<值类型，
type persons struct { //父结构体
	Name    string
	Age     int    `json:"age"` //结构体表情-转换 json 后重命名
	sex     string //私有属性不能被 json 包访问
	address        //嵌套匿名结构体
}
type address struct {
	city string
	time int
}
type women struct { //子结构体
	sleep   string
	persons //结构体嵌套，继承父
}

func (p *persons) personGo(name string, age int) { //结构体方法
	p.Name = name
	p.Age = age
}
func init() {
	fmt.Println("main init...")
}

//接口
type Usber interface {
	start()
	// Ainterface //接口嵌套
}
type Computer struct{}
type Phone struct {
	Name string
}

func (p Phone) start() { //接口地址 *phone 用来修改结构体
	fmt.Println(p.Name, "start....")
}
func (c Computer) work(usb Usber) {
	usb.start()
}
func main() {
	// 字符串<<<<<<<<<<<<<<<
	str := "123-456-789" //全局变量不能用类型推导
	arr := strings.Split(str, "-")
	str1 := strings.Join(arr, "=")
	fmt.Println(arr, str1, len(str), strings.Index(str, "5"), strings.HasPrefix(str, "123"))
	fmt.Printf("值:%v 原样:%c 类型:%T\n", str[1], str[1], str[1])
	// Sprintf 转换：%d int %.2f float %t bool %c byte |  也可以用 strconv 转换  strconv.FormatInt(int64(X),10) | num,err=strconv.ParseInt(str,10,64) string 转 int
	str2 := []byte(str) //rune 中文  字符串无法修改，只能转换
	str2[0] = '0'       //字符
	fmt.Println(string(str2))
	// 运算符<<<<<<<<<<<<<<<     除法如果都是整数，结果也是整数, ++和--只能独立使用
	// 流程控制<<<<<<<<<<<<<<<
lable2:
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i == 2 {
			break lable2 //多重循环用
		} else {
			goto lable3 //无条件跳转
		}
		// continue 结束当前循环，开始下一次循环
	}
	for key, val := range str {
		fmt.Printf("key:%v value:%c\n", key, val)
	}
lable3:
	var vari = 8
	switch vari {
	case 1, 2, 3:
		fmt.Println("xiao")
		fallthrough //不需要 break， fallthrough 只穿透下一层
	case 4, 5:
		fmt.Println("345")
	}
	// 数组<<<<<<<<<<<<<<<数组长度也是数组的一部分,长度不可变，基本数据类型和数组都是值类型
	var arr2 = [3]string{"js", "python", "java"}
	arr3 := [...]int{1: 4, 2: 2} //...推断长度，|下标来创建
	arr4 := [2][2]string{        //二维数组,多维数组只有第一个支持...推导长度
		{"app", "andriod"},
		{"ios", "php"},
	}
	arr4[0][1] = "andriods"
	fmt.Println(arr2, arr3, arr4)
	// 切片<<<<<<<<<<<<<<引用数据类型,nil
	slice6 := make([]int, 3, 3)
	slice1 := []int{7, 2, 3}
	slice2 := slice1[:1] //截取 X：Y,基于数组截取也是切片,不截取 Y
	fmt.Println("slice---", slice1, slice2, slice6)
	fmt.Printf("长度：%v 容量:%v\n", len(slice2), cap(slice2)) //容量，从他第一个元素开始，到底层数组元素末尾的个数
	slice1 = append(slice1, slice2...)                    //追加slice/合并
	slice3 := make([]int, 4, 6)
	copy(slice3, slice1) //不影响原切片
	slice3[0] = 9
	slice4 := append(slice1[:1], slice1[2:]...)    //只能这样删除
	sort.Ints(slice1)                              //sort.Float64s()    sort.Strings()
	sort.Sort(sort.Reverse(sort.IntSlice(slice1))) //降序   Float64Slice   StringSlice
	fmt.Println(slice1, slice3, slice4)
	// map<<<<<<<<<<<<<<引用数据类型,nil   for range 循环
	userInfo := map[string]string{
		"name": "猴子",
		"age":  "12",
		"sex":  "男",
	}
	v, ok := userInfo["age"]
	delete(userInfo, "sex")
	fmt.Println(userInfo, v, ok)
	fmt.Println(func1(5, 6, 7))
	//函数《《《《《《《《《《《
	su, sb := func2(7, 8)
	fmt.Println(su, sb)
	sum, sub := calc(7, 9, func2) //fun2也可以换匿名函数
	fmt.Println(sum, sub, f3())
	fmt.Println(f4(10, 0))
	//Time《《《《《《《《《《《 2006年01月02日 03|15时 04分05秒
	timeObj := time.Now() // timeObj,_ := time.parseInLocation("2006...",timeString,time.Local)字符串转时间戳 | 时间戳转换成 OBJ，time.Unix(int64(unixtime),0)
	ticker := time.NewTicker(time.Second / 3)
	times := 3
	for t := range ticker.C {
		if times == 0 {
			ticker.Stop()
			break
		}
		times--
		fmt.Println(t)
	}
	time.Sleep(time.Second / 3)
	fmt.Println(timeObj, timeObj.Format("2006-01-02 03:04:05"), timeObj.Unix())
	//指针<<<<<<<<<<<<也是一种特殊的变量,引用数据类型，存储的不是普通值，二是另一个变量的内存地址
	pointA := 10
	pointB := &pointA
	*pointB = 30
	fmt.Printf("指针值-%v 类型:%T,地址:%p,A-%v ", pointB, pointB, &pointB, pointA)
	pointFun(&pointA)
	fmt.Println("指针改变值", pointA, *pointB)
	//struct 结构体<<<<<<<<<<<<值类型,结构体实例是独立的
	person1 := persons{
		Name: "张三", //可以都没有 key
		Age:  13,
		sex:  "男",
	} //var person1 persons | var person1 = new(persons) 获得的是指针，结构体内也支持 person1.name 获取  | var person1 = &persons{} 也是指针
	// person1.name = "张三"
	// person1.age = 17
	// person1.sex = "男"
	person1.personGo("王五", 14)
	fmt.Printf("值：%v 类型:%T \n 所有:%#v", person1, person1, person1)
	person1.city = "深圳"
	person1.address.time = 12
	fmt.Println(person1)
	//结构体继承
	women1 := women{
		sleep: "睡觉",
		persons: persons{ //也可以传结构体指针
			Name: "小花",
			Age:  14,
		},
	}
	women1.personGo("小花花", 15)
	fmt.Printf("women1 : %#v\n", women1)
	//json 转换   结构体转换成 json 字符串
	jsonBype, _ := json.Marshal(women1) //byte 类型的切片
	fmt.Println("json", jsonBype, "\n", string(jsonBype))
	//字符串转换成结构体
	jsonString := `{"Name":"小花花json","Age":15}`
	personJson := persons{}
	err := json.Unmarshal([]byte(jsonString), &personJson)
	fmt.Println(err, personJson)
	//自定义包
	sum5 := calcs.Add(9, 7)
	fmt.Println(sum5)
	//三方包decimal 处理 float 运算精度丢失
	float1 := 3.1
	float2 := 4.2
	fmt.Println(float1+float2, decimal.NewFromFloat(float1).Add(decimal.NewFromFloat(float2))) //减法 Sub 乘法 Mul 除法 Div
	//接口也是一种数据类型,由具体类型和类型的值两部分组成		|   定义了对象的行为规范，只定义不实现
	var Computer1 = Computer{}
	var phone = Phone{Name: "一加"}
	Computer1.work(phone)
	//空接口表示没有任何约束
	//  空接口不支持索引   比如X为空接口类型 X.name / X[1]  处理： Y,_ := X.([]string)  ==>  Y.name     []string 这是 X 的类型
	var nilInterface interface{}
	nilInterface = 3
	nilInterface = "null"
	fmt.Println("空接口可以是任意类型", nilInterface)
	//类型断言
	val2, OKKK := nilInterface.(string)
	fmt.Println(val2, OKKK)
}
