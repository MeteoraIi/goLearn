package main

import "fmt"

func main(){
	arr := []int{1, 2, 3}
	sum := 0

	for i, num := range arr{
		sum += num
		if num == 2{
			fmt.Println("index:", i, ", num: ", num)
		}
	}
	fmt.Println("sum:", sum)

	m := map[string]string{"a": "A", "b": "B"}

	for k, v := range m{
		fmt.Println("key:", k, ",value:", v)
	}

	for k := range m{
		fmt.Println("key:", k)
	}
}