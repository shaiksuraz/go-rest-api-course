package main

import "fmt"

func Run() error {
	fmt.Println("starting our application")
	return nil
}

func main() {
	fmt.Println("GO REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
