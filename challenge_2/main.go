package main

import (
	"fmt"
)

var val = "initial global"

func main() {
	val := 12
	var pval *int = &val

	fmt.Printf("Local val: %v (%T)\n", val, val)
	printGlobal()
	updateGlobal("updated global")
	printGlobal()
	fmt.Printf("Pointer to local val: %v (%T)\n", pval, pval)
	fmt.Printf("Dereferenced pointer to local val: %v (%T)\n", *pval, *pval)

	*pval = 15
	fmt.Printf("Reassigned local val via pointer: %v (%T)\n", *pval, *pval)

	// or can do pointer things without creating a separate pointer var
	fmt.Printf("Address of local val: %v (%T)\n", &val, &val)
	*(&val) = 20 // why?
	fmt.Printf("Reassigned local val via black magic: %v (%T)\n", val, val)
}

func printGlobal() {
	fmt.Printf("Global val: %v (%T)\n", val, val)
}

func updateGlobal(newValue string) {
	val = newValue
}
