package main

import "fmt"

func detectType(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("Тип: int, значение = %d\n", v)
	case string:
		fmt.Printf("Тип: string, значение = %q\n", v)
	case bool:
		fmt.Printf("Тип: bool, значение = %v\n", v)
	case chan int:
		fmt.Println("Тип: chan int")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	detectType(42)
	detectType("hello")
	detectType(true)
	detectType(make(chan int))
}
