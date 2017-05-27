package main

import "fmt"

func main() {

	done1 := make(chan int)
	done2 := make(chan int)
	im :=  []int{1, 2, 3, 4, 5}

	for i := range im {
		defer func() {
			fmt.Print("defer A ")
			fmt.Println(i) //тут ошибка - захвачена итерационная переменная
		}()
	}
	fmt.Println("defer A started")

	for i := range im {
		defer func(n int) {
			fmt.Print("defer B ")
			fmt.Println(n)
		}(i)
	}
	fmt.Println("defer B started")

	for i := range im {
		go func() {
			fmt.Printf("go A %v\n",i) //тут ошибка - захвачена итерационная переменная
			done1 <- 1
		}()
	}
	fmt.Println("go A started")

	for range im {<-done1}

	for i := range im {
		go func(n int) {
			fmt.Printf("go B %v\n",n)
			done2 <- 1
		}(i)
	}
	fmt.Println("go B started")

	for range im {<-done2}

}
