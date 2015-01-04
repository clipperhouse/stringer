// Run this before testing: go setup.go

package main

import (
	"fmt"

	_ "github.com/clipperhouse/stringer"
	"github.com/clipperhouse/typewriter"
)

func main() {
	a, err := typewriter.NewApp("+test")
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := a.WriteAll(); err != nil {
		fmt.Println(err)
		return
	}
}
