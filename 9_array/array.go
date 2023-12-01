package main

import "fmt"

func main() {
	
	var checks [3]bool
	fmt.Printf("checks: %#v\n", checks)

	var skills = [...]string{"js", "go", "python"}
	show(skills)
	skills[1] = "golang"
	fmt.Printf("skills: %#v\n", skills)
	fmt.Printf("skills: %#v\n", skills[1])
	fmt.Printf("skills: %#v\n", len(skills))
}

func show(sk [3]string) {
	fmt.Printf("show: %#v\n", sk)
}