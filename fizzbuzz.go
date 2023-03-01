package main
import(
	"fmt"
)
func main(){
	for i:= 1; i<=20;i++{
		if i%3 == 0 && i%5 == 0{
	fmt.Printf("%v fizz buzz\n",i)
	}else if i%3==0{
		fmt.Printf("%v buzz\n",i)

	}else if i%5==0{
		fmt.Printf("%v fizz\n",i)
	}else{
		fmt.Printf("%v\n",i)
	}


	}
	check := []int{1,2}
	for i,j := range check {
		fmt.Println(i,j)
	}
}