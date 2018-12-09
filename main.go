package main

import (
	"fmt"
	"gotuts/exercise/dog"
)

func main() {
	humanYear := 12
	dogYear, err := dog.Convert(humanYear)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dogYear)
}
