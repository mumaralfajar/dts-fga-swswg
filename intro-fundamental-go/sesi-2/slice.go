package main

// import "fmt"

// func main() {
// 	var marvelHeroes []string = []string{
// 		"Colossus",
// 		"Sun Fire",
// 		"Wolverine",
// 		"Juggernaut",
// 		"Ice Man",
// 		"Magneto",
// 	}

// 	var randomHeroes []string = marvelHeroes[2:5]

// 	// fmt.Printf("marvelHeroes: %#v \n", marvelHeroes)
// 	// fmt.Printf("randomHeroes: %#v \n", randomHeroes)

// 	// fmt.Printf("Marvel Heroes (length: %d, capacity: %d) \n", len(marvelHeroes), cap(marvelHeroes))
// 	// fmt.Printf("Random Heroes (length: %d, capacity: %d) \n", len(randomHeroes), cap(randomHeroes))

// 	// & => ampersand

// 	randomHeroes = append(randomHeroes, "Wonder Woman")

// 	// fmt.Printf("Random Heroes (length: %d, capacity: %d) \n", len(randomHeroes), cap(randomHeroes))
// 	// fmt.Println("Marvel Heroes index[2]:", &marvelHeroes[2])
// 	// fmt.Println("Random Heroes index[0]:", &randomHeroes[0])

// 	randomHeroes = append(randomHeroes, "Sabertooth")

// 	// fmt.Printf("Random Heroes (length: %d, capacity: %d) \n", len(randomHeroes), cap(randomHeroes))

// 	// fmt.Println("Marvel Heroes index[2]:", &marvelHeroes[2])
// 	// fmt.Println("Random Heroes index[0]:", &randomHeroes[0])

// 	// fmt.Println("Length of marvelHeroes:", len(marvelHeroes))
// 	// fmt.Println("Capacity of marvelHeroes:", cap(marvelHeroes))

// 	_ = marvelHeroes

// 	var a [4]int = [4]int{10, 20, 30, 40}

// 	var b []int = a[:]

// 	fmt.Println("Cap of b:", cap(b))

// 	b = append(b, 500) //len => 5, cap => 8

// 	b[0] = 120

// 	fmt.Printf("a %#v \n", a)
// 	fmt.Printf("b %#v \n", b)

// 	_ = b
// }
