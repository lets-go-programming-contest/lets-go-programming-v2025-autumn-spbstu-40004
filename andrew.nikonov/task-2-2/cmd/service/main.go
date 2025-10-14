package main

import (
	"errors"
	"fmt"
)

var errInput = errors.New("incorrect input")

func main() {
	var dishesNum int

	_, err := fmt.Scanln(&dishesNum)
	if err != nil || dishesNum < -10000 || dishesNum > 10000 {
		fmt.Println(errInput)

		return
	}

	//TO DO: reading and processing heap

	var preferredNum int

	_, err = fmt.Scanln(&preferredNum)
	if err != nil || preferredNum < 1 || preferredNum > dishesNum {
		fmt.Println(errInput)

		return
	}

}
