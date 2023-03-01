package main

import (
	//"bufio"
	"fmt"
	//"sort"
	//"math"
	//"os"
	//"strconv"
//	"strings"
)


func main() {
	number:=make(map[string]string)
	number["sai"]="shashank"
	fmt.Print(number["sai"])
	for k,v := range(number){
		fmt.Printf("\n%v %v",k,v)
	}


}
