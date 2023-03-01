package main
import(
"fmt"
"strings"
)
func main(){
	input1:="this is a string of type a and string b of type b"
	maps:= map[string]int{}
	word:=strings.Fields(input1)
	for _,j := range word {
		maps[strings.ToLower(j)]++



	}
	fmt.Println(maps)


}