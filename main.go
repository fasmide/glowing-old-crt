package main

import (
	"log"
)

func main() {
	log.Println("This is go")
	v := NewVideo("https://www.youtube.com/watch?v=C_3d6GntKbk")
	v.Download()

	player := NewMPlayer()

	player.Play(&v.Data)

	player.Play(&v.Data)

}
