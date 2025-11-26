package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	x := rand.Intn(100)
	scanner := bufio.NewScanner(os.Stdin)
el:
	for i := 0; i < 10; i++ {
		fmt.Println("Введите число:")
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		if x == y {
			fmt.Printf("Число угадано количество попыток: %d", i)
			break el
		} else if x > y {
			fmt.Println("Заданное число больше ")
		} else {
			fmt.Println("Заданное число меньше ")
		}
	}
	fmt.Println("Число не угаданно")
}
