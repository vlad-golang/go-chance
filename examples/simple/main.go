package main

import (
	"fmt"

	"github.com/vlad-golang/go-chance"
)

func main() {
	if chance.Percentage(30) {
		fmt.Println("Событие с вероятностью 30% произошло!")
	}
}
