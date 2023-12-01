package main

import "fmt"

func main() {

	// Slice dynamic lenge of array

	skills := []string{"js", "go", "python"}
	fmt.Printf("%T => skills: %v\n", skills, skills)

	skills[1] = "golang"
	fmt.Printf("%#v\n", skills)

	s := skills[0]
	fmt.Printf("%#v\n", s)

	for i := 0; i < len(skills); i++ {
		fmt.Printf("for => %#v\n", skills[i])
	}

	for _, val := range skills {
		fmt.Printf("for range => %#v\n", val)
	}

	//append
	fmt.Printf("%p\n", &skills)
	skills = append(skills, "ruby", "java", "erlang")
	fmt.Printf("%p\n", &skills)
	show("skills", skills)

	genres := []string{"Action", "Adventure", "Fantasy", "Sci-Fi"}
	show("genres", genres)

	s1 := skills[0:2] //Math: Half-Open range
	show("s1", s1)

	show("skills[1:4]", skills[1:4])

	show("skills[:]", skills[:]) // not specific first and last index

	var ss []int
	fmt.Printf("length: %d => %#v\n", len(ss), ss)
	if ss == nil {
		fmt.Println("ss is nil")
	}
	ss = append(ss, 33)
	fmt.Printf("length: %d => %#v\n", len(ss), ss)

	vv := make([]int, 3) // make initial slice 3 items
	fmt.Printf("length: %d => %#v\n", len(vv), vv)

	//overlap
	skills2 := []string{"js", "go", "python"}
	x := skills2[0:2]
	y := skills2[1:3]
	show("skills2", skills2)
	show("x", x)
	show("y", y)

	fmt.Println("===change value by reference===")
	x[1] = "xxx"
	show("skills2", skills2)
	show("x", x)
	show("y", y)
}

func show(text string, slice []string) {
	fmt.Printf("%s : length[%d],capacity[%d] => %v\n", text, len(slice), cap(slice), slice)
}
