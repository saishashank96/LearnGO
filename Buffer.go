package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//take input using buffere convert to float and add

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Number:")
	input1, _ := reader.ReadString('\n')
	fmt.Println("Number is :", input1)
	fmt.Print("Enter Number:")
	input2, _ := reader.ReadString('\n')
	fmt.Println("Number is :", input2)
	input3,_:=strconv.ParseFloat(strings.TrimSpace(input1),64)
	input4,_:=strconv.ParseFloat(strings.TrimSpace(input2),64)	
	fmt.Println(add(input3,input4))


}
func add(a float64,b float64) float64 {
	return a + b
}
