package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Andre","Last":"Sampaio","Age":25},{"First":"Eliete","Last":"Sampaio","Age":49}]`
	bs := []byte(s)

	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)
}
