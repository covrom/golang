package main
import (
	"fmt"
	"os"
	"strings"
)

//коммент
func main(){
	// fmt.Println("1+1=",1+1)
	// fmt.Println(len("Привет всем"))
	// fmt.Println("Привет всем"[12]) //берет побайтно, юникодовые символы считает по два байта, остальные по одному
	// fmt.Println("Привет"+"всем")
	// fmt.Println(true && true)
	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(true || true)
	// fmt.Println(!true)

	// var x string 
	// var x = "Привет всем"

	// x := "Привет всем"
	// x = x+"!"
	// fmt.Println(x)

	// fmt.Print("Дай число:")
	// var num float64
	// fmt.Scanf("%f",&num)
	// dbl := num*num
	// fmt.Println(dbl)

	// fmt.Println(`ikgh
	// piohoijh
	// ;ih
	// `)

	// i:=1
	// for i<=10{
	// 	if i%2==0{
	// 		fmt.Println(i,"четное")
	// 	}else{
	// 		fmt.Println(i,"нечетное")
	// 	}
		
	// 	i++
	// }

	// elems := map[string]string{}

	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:]," "))


}