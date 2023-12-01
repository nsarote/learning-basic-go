package main

import "fmt"

func main() {
	fmt.Println("\n===[summation]===")
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println("i: ", i, " sum: ", sum)
	}
	fmt.Println("sum done: ", sum)


	fmt.Println("\n===[summation until]===")
	sum = 1
	for ; sum < 5;  {
		sum += sum
		fmt.Println("sum: ", sum)
	}
	fmt.Println("sum done: ", sum)

	
	fmt.Println("\n===[while]===")
	sum = 1
	for sum < 5 { // like while go have only for
		sum += sum
		fmt.Println("sum: ", sum)
	}
	fmt.Println("sum done: ", sum)


	fmt.Println("\n===[iterate array]===")
	var skills = [...]string{"js", "go", "python"}
	for i := 0; i < len(skills); i++ {
		fmt.Println("skills[",i,"]: ", skills[i])
	}

	fmt.Println("\n===[iterate array using range]===")
	for i := range skills {
		fmt.Println("skills[",i,"]: ", skills[i])
	}

	fmt.Println("\n===[iterate array using range and val]===")
	for i, val := range skills {
		fmt.Println("skills[",i,"]: ", val)
	}

	fmt.Println("\n===[iterate array using range and use only val]===")
	for _, val := range skills {
		fmt.Println("value: ", val)
	}
}
