package main

import "fmt"

// Using the code from the previous example, add a record to your map. Now print the map out
// using the “range” loop

func main() {

	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	m["sampaio"] = []string{"Hi", "Mom"}

	for _, v := range m["sampaio"] {
		fmt.Println(v)
	}

}
