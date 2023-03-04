package main
import(
"fmt"
)
const (isadmin = 1 << iota
 iseng
 india
 usa
china
)
func main(){
	const person  = isadmin | india
	fmt.Println(((person & isadmin)==isadmin) && ((person & india)==india))
}