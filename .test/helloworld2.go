package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    for i:=0; i < 1; i = i+1 {
       fmt.Printf("%s\n",  strings.ToUpper(os.Args[1]))
    }
}
