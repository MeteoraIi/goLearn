package main

import (
	"fmt"
	// "reflect"
)

func main() {
	// var sArr = "fuckMe"

	// s1 := sArr[0:3]
	s2 := make([]string, 3)
	s2[0] = "a"
	s2[1] = "b"
	s2[2] = "c"


	fmt.Println("get:", s2[1])
	// fmt.Println("len s1:", len(s1))
	fmt.Println("len s2:", len(s2))
	fmt.Println("cap s2:", cap(s2))

	s2 = append(s2, "d")
	s2 = append(s2, "e", "f")
	fmt.Println(s2)

	c := make([]string, len(s2))
	copy(c, s2)
	fmt.Println(c)

	fmt.Println(s2[2:5])
	fmt.Println(s2[:5])
	fmt.Println(s2[2:])

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good)

	// fmt.Println(reflect.TypeOf(s1))
	// fmt.Println(reflect.TypeOf(s2))

}
