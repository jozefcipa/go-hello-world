package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, _ := user.Current()

	fmt.Printf("Hello %s!\n", user.Name)
}
