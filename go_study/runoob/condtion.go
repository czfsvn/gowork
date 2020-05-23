package main

import "fmt"

func if_condition() {
	a := 10
	if a < 10 {
		fmt.Println("if_condition: low")
	} else if a == 10 {
		fmt.Println("if_condition: equal")
	} else {
		fmt.Println("if_condition: greater")
	}
}

/*
switch: 匹配项后面也不需要再加 break。
	默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，
	如果我们需要执行后面的 case，可以使用 fallthrough
*/
func switch_condition1() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 60

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

func switch_condition2() {
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}

/*
select: 用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
	随机执行一个可运行的 case。
	如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。
*/
func select_condition() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default: /* 可选 */
		fmt.Printf("no communication\n")
	}
}

func condtion_test() {
	if_condition()
	switch_condition1()
	switch_condition2()
	select_condition()
}

func main() {
	condtion_test()
}
