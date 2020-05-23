package main

import (
	"fmt"
	"math"
)

/*
	func function_name( [parameter list] ) [return_types] {
   		函数体
	}
*/

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/*
值传递：传递是指在调用函数时将实际参数复制一份传递到函数中，
	这样在函数中如果对参数进行修改，将不会影响到实际参数。
*/
func swap(x, y string) (string, string) {
	return y, x
}

/*
值传递是 指在调用函数时将实际参数复制一份传递到函数中，
这样在函数中如果对参数进行修改，将不会影响到实际参数。
*/
func swap_value(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}

/*
引用传递是 指在调用函数时将实际参数的地址传递到函数中，
那么在函数中对参数所进行的修改，将影响到实际参数。
*/
func swap_ref(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

func func_pass_value() {
	fmt.Println("func_pass_value------------------")
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前，a/b 的值 : %d/%d\n", a, b)

	/* 通过调用函数来交换值 */
	swap_value(a, b)

	fmt.Printf("交换后，a/b 的值 : %d/%d\n", a, b)
}

func func_pass_ref() {
	fmt.Println("func_pass_ref------------------")
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前，a/b 的值 : %d/%d\n", a, b)

	/* 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	swap_ref(&a, &b)

	fmt.Printf("交换后，a/b 的值 : %d/%d\n", a, b)
}

func func_call() {
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
}

func func_main() {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
	func_pass_value()
	func_pass_ref()
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 闭包使用方法
func add(x1, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x1 * x2
	}
}

func func_closure() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	add_func := add(1, 2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())

	add_func2 := add(3, 2)
	fmt.Println(add_func2())
	fmt.Println(add_func2())
	fmt.Println(add_func2())
	fmt.Println(add_func2())
	fmt.Println(add_func2())
}

/*
方法：
	func (variable_name variable_data_type) function_name() [return_type]{
   		// 函数体
	}
	C++ 等语言中，实现类的方法做法都是编译器隐式的给函数加一个 this 指针，
	而在 Go 里，这个 this 指针需要明确的申明出来，其实和其它 OO 语言并没有很大的区别。
*/

/* 定义结构体 */
type Circle struct {
	radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func method() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

func main() {
	//func_main()
	//func_call()
	//func_closure()
	method()
}
