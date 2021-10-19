// Learning & illustrating different ways to use print statements
// And, declaring variables.

package go_learning

import (
	"fmt"
	"go_learning/calling_an_ext_mod"
)

func main() {
	variable_declarations()
	fmt.Print(`testing
multi
line
strings in
goLang`)
	calling_an_ext_mod.Ext_mod()
}

/*
Multi line comments
*/

func variable_declarations() {
	const hello = "akjsdk"
	fmt.Print("A constant buddy  ", hello)
	name := "a simple variable"
	fmt.Print("this is another way to declare a var ", name)
	var namee = "hey "
	fmt.Printf("%s this is one more way to declare a var ", namee)
	var integ int = 34
	fmt.Println(integ)
}
