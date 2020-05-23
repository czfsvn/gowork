package main // package name

import "fmt" // import package name
/*
变量声明1:指定变量类型，如果没有初始化，则变量默认为零值
零值就是变量没有做初始化时系统默认设置的值
*/
func para_declare1() {

	fmt.Print("function para_declare1\n")
	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)
}

/*
变量声明2:根据值自行判定变量类型
*/
func para_declare2() {
	fmt.Print("function para_declare2\n")
	var d = true
	fmt.Println(d)
}

/*
变量声明3:	省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误
			但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。
*/
func para_declare3() {
	fmt.Print("function para_declare3\n")
	d := 10
	s := "para_declare3"
	fmt.Println(d, s)
}

/*
多变量声明
*/
func para_declare4() {
	fmt.Print("function para_declare4\n")
	var d, s, f = 123, "para_declare4", 0.56
	fmt.Println("para_declare4", d, s, f)

	var a, b int
	var c string
	a, b, c = 5, 7, "abc" //并行 或 同时 赋值。
	_, glob := 5, 7       //5 被抛弃了，空白标识符 _
	fmt.Println("para_declare4", a, b, c, glob)
}

//一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}

/*
从函数返回多个值
*/
func para_declare5() {
	fmt.Print("function para_declare5\n")

	a, b, c := numbers()
	fmt.Println("para_declare5", a, b, c)

	d, _, c := numbers()
	fmt.Println("para_declare5", d, c)
}

func para_declare() {
	//para_declare1()
	//para_declare2()
	//para_declare3()
	para_declare4()
	para_declare5()
}

func main() { // 函数的花括号不允许换行
	para_declare()
}
