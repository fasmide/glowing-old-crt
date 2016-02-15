package main

import (
	"log"
)

func main() {
	log.Println("This is go")
	v := NewVideo("https://www.youtube.com/watch?v=jAZZaslGxrk")
	v.Download()

}
