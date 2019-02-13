package main

import "fmt"
import "time"
import "math/rand"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Your random number is", rand.Intn(500))
}