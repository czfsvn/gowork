package main

import "fmt"

/*
	1、for init; condition; post { }	与c语言for一样
	2、for condition { }				与C语言while一样
	3、for { }							与c语言 for(;;)
*/
func for_loop1() {
	sum := 0
	for i := 0; i <= 50; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	sum = 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)
}

func for_loop2() {
	/* 定义局部变量 */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

func for_loop4() {
	for m := 1; m < 10; m++ {
		/*    fmt.Printf("第%d次：\n",m) */
		for n := 1; n <= m; n++ {
			//fmt.Printf("%dx%d=%d ", n, m, m*n)
			fmt.Printf("%d+%d=%d ", n, m, m+n)
		}
		fmt.Println("")
	}
}

func for_loop3() {
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

func for_loop() {
	for_loop1()
	for_loop2()
	for_loop3()
	for_loop4()
}

/*
第一步：假设苹果和橘子各少卖300千克，即苹果没有卖出，就意味着橘子卖掉1500千克，
	这时橘子和苹果的重量相等。

第二步：由于之前橘子是苹果的3倍，而现在橘子卖掉1500千克后与苹果重量相等

第三步：所以 1500 千克刚好是苹果的 2 倍
*/
func break_loop1() {
	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}
}

//在多重循环中，可以用标号 label 标出想 break 的循环。
func break_loop2() {
	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re //执行1-11， 然后跳到re 执行2-11， 再跳到re 执行 3-11
		}
	}
}

func break_loop() {
	break_loop1()
	break_loop2()
}

func continue_loop1() {
	// 不使用标记
	fmt.Println("---- continue ---- ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue
		}
	}
}

func continue_loop2() {
	// 使用标记
	fmt.Println("---- continue label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue re //执行1-11， 然后跳到re 执行2-11， 再跳到re 执行 3-11
		}
	}
}

func continue_loop() {
	continue_loop1()
	continue_loop2()
}

func goto_loop() {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}

func main() {
	//for_loop()
	//for_break()
	//continue_loop()
	goto_loop()
}
