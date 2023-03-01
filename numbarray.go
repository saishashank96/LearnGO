package main

import (
	//"bufio"
	"fmt"
	"sort"
	//"math"
	//"os"
	//"strconv"
//	"strings"
)


func main() {
	number:=make([]int,5)
	number[0]=1
	number=append(number,3)
	sort.Sort(sort.Reverse(sort.IntSlice(number)))
	fmt.Print(number)


}
