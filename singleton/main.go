package main

import (
	"fmt"

	"github.com/gitfuf/misc/singleton/theone"
)

func main() {
	s := theone.New("the only one")
	print(s)

	s2 := theone.New("the second instance ...oops... it is singleton")
	print(s2)
}

func print(s theone.TheOne) {
	fmt.Println("Singleton: ", s)
}
