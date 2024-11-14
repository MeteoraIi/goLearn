package main

import "fmt"

type user struct{
	name string
	password string
}

func main(){
	var a user = user{name: "wang", password: "1024"}
	b := user{"wu", "2048"}
	c := user{name: "sun"}
	c.password = "4096"

	var d user
	d.name = "li"
	d.password = "8078"

	fmt.Println(a, b, c, d)
	fmt.Println(checkPassword(a, "1024"))
	fmt.Println(checkPassword2(&d, "2323"))
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}