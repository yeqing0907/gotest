package main

import (
	"fmt"
	"strings"
)

var isActive bool
var a bool = true
var b int = 20

const c1 string = "abc"
const (
	Unknown = 0
	Female = 1
	Male = 2
)

#main is the enter
func main() {

	var i = 1.5
	var j = 2
	fmt.Println(i, j)

	stu := "studbets"

	j2 := "pass"
	fmt.Println(j2, stu)

	fmt.Println(isActive)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(c1)

	fmt.Println(Unknown)
	fmt.Println(Female)
	fmt.Println(Male)

	fmt.Println("hello world")

	str := "这里是 www\n.runoob\n.com"
	fmt.Println("-------- 原字符串 ----------")
	fmt.Println(str)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println("-------- 去除空格与换行后 ----------")
	fmt.Println(str)


	/* 定义局部变量 */
	var a2 int = 100

	/* 使用 if 语句判断布尔表达式 */
	if a2 < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a2 小于 20\n" )
	}else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a2 不小于 20\n" );
	}
	fmt.Printf("a2 的值为 : %d\n", a2)


	var cnt = 5;
	for i := 0; i < cnt; i++ {
		fmt.Println(i);
	}

}
