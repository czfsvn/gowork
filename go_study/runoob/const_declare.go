package main

import (
	"fmt"
	"unsafe"
)

func const_declare_1() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d\n", area)
	//println()
	println("const_declare_1", a, b, c)
}

func const_declare_2() {
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)

	fmt.Println("const_declare_2", Unknown, Female, Male)
	fmt.Println("const_declare_2", a, b, c)
}

func const_declare_iota() {
	/*
		iota，特殊常量，可以认为是一个可以被编译器修改的常量。
		iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
			const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

		iota 可以被用作枚举值：
			第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
	*/
	const (
		x = iota
		y = iota
		z = iota
	)
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)

	const (
		m = 1 << iota // 1 << 0
		j = 3 << iota // 3 << 1
		k             // 3 << 2
		l             // 3 << 3
	)
	fmt.Println("const_declare_iota", x, y, z)
	fmt.Println("const_declare_iota", a, b, c, d, e, f, g, h, i)
	fmt.Println("const_declare_iota", m, j, k, l) // ? 没有明白
}

func test_main() {
	//const_declare_1()
	//const_declare_2()
	const_declare_iota()
}

func main() {
	test_main()
	fmt.Println("const declare")
}
