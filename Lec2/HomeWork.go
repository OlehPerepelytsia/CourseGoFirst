package main

import "fmt"

func main() {
	fmt.Println("Привет! Я новый студент!")
	var (
		world1 = "имя"
		world2 = "твое"
		world3 = "мне"
		world4 = "знакомо"
	)
	fmt.Printf("%s %s %s %s\n%s %s %s %s\n%s %s %s %s\n", world4, world3, world2, world1, world3, world1, world4, world2, world2, world4, world1, world3)
}
