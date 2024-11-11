package main

import (
	"fmt"

	"github.com/AEKDA/aebuilder/internal/cli"
)

func main() {
	cmd := cli.New()
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Done!")
}
