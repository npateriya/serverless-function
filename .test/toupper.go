package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    for i := range os.Args{
       if i != 0 {
       		fmt.Printf("%s\n",  strings.ToUpper(os.Args[i]))
      }
    }
}
