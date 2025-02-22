
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hello, World Wide"

	fmt.Println(strings.Contains(greeting, "Hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Booni"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))

	//The Original value is unchange
	fmt.Println("Original string value =", greeting)

	// Sort Package
	ages := []int{12, 30, 56, 78, 89, 56, 23, 45}
	names := []string{"sam", "dare", "afolabi", "seyi"}

	sort.Ints(ages)
	sort.Strings(names)
	fmt.Println(ages)

	table := sort.SearchStrings(names, "dare")
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)
	fmt.Println(table)

}
