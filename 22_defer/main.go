package main

import "fmt"

func main() {
	defer fmt.Println("world 1") // => ทำภายหลังก่อนจบ func เป็น stack
	defer fmt.Println("world 2")
	defer fmt.Println("world 3")
	fmt.Println("hello")
	test()
	test2()
}

func test() {
	fmt.Println("start test => counting")
	for i:=0;i<10;i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}

func test2() {
	fmt.Println("start test => counting")
	for i:=0;i<10;i++ {
		/*defer func() {
			fmt.Println(i) // WARNING Scope of defer
		}()*/
		defer func(n int) {
			fmt.Println(n) // WARNING Fix above issue by pass by value to defer func
		}(i)
	}
	fmt.Println("Done")
}
