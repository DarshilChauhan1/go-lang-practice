package main

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		println(j)
		j++
	}

	// infinite loop
	// for {
	// 	println("infinite loop")
	// }

	// loop through a collection
	names := []string{"Gladys", "Samantha", "Darrin"}
	for k, v := range names {
		println(k, v)
	}
}