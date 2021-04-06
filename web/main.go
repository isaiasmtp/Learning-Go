package main

import "fmt"

func main() {
	ok, err := say("Hello Word")

	if err != nil {
		panic(err.Error())
	}
	switch ok {
	case true:
		fmt.Println("Worked")
	case false:
		fmt.Println("Dont Worked")
	}
}

func say(what string) (bool, error) {
	if what == "" {
		return false, fmt.Errorf("Empty String")
	}

	fmt.Println(what)

	return true, nil
}
