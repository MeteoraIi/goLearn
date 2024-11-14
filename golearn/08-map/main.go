package main

import (
	"fmt"
)

func main(){
	var m = make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["one"])
	fmt.Println(m["unknow"])

	v, e := m["unknow"]
	fmt.Println(v, e)

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = m2
	fmt.Println(m2, m3)
}