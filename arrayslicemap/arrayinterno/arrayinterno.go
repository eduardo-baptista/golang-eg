package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)

	s2 := append(s1, 1, 2, 3) // ao utlizar append é mantido o msm array de referencia

	fmt.Println(s1, s2)

	s1[0] = 7 // alteração se reflete em ambos os slices

	fmt.Println(s1, s2)

}
