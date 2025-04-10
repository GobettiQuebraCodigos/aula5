package main

import (
	"fmt"
)

func main() {
	age := 16
	fmt.Println(age <= 20)
	fmt.Println(age >= 20)
	fmt.Println(age == 20)
	fmt.Println(age != 20)

	if age < 30 {
		fmt.Println("Menor que 30 anos")
	} else {
		fmt.Println("Maior que 30 anos")
	}

	names := []string{"Lucas", "Leonardo", "Rafael"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("Continue após a posição", index, "e valor", value)
			continue
		}
		if index > 2 {
			fmt.Println("Sair após", index)
			break
		}
		fmt.Println("Valor: ", value)
	}
}