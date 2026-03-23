package main

import (
	"fmt"
	"os"
)

func main() {
	Echo1(os.Args)
}

func Echo1(args []string) {
	for i := 1; i < len(args); i++ {
		fmt.Printf("%s ", args[i])
	}
	fmt.Println()
}
