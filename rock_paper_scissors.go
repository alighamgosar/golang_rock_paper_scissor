package main

import (
    "fmt"
    "math/rand"
	"reflect"
)



func main() {

	A_score := 0
	B_score := 0

	wins := map[string]string{"ROCK": "SCISSOR", 
						  "PAPER": "ROCK", 
						  "SCISSOR": "PAPER"}
	choices := reflect.ValueOf(wins).MapKeys()

	i := 1
	for (A_score < 3 && B_score < 3) {

		fmt.Println("Iteration number :", i)
		i += 1
		

		index1 := rand.Intn(len(choices))
		index2 := rand.Intn(len(choices))

		A_choice := choices[index1].String()
		B_choice := choices[index2].String()


		fmt.Println("Player A selected: ", A_choice)
		fmt.Println("Player B selected: ", B_choice)

		if (A_choice == B_choice){
			fmt.Println("Draw!")
		} else if A_choice == wins[B_choice] {
			fmt.Println("B Won!")
			B_score += 1
		} else {
			fmt.Println("A won!")
			A_score += 1
		}
		fmt.Println("\n")

		
	}

	fmt.Printf("Result: %d - %d", A_score, B_score)



}
