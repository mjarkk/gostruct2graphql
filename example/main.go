package main

import (
	"fmt"
	"os"

	"github.com/mjarkk/gostruct2graphql"
	"github.com/mjarkk/gostruct2graphql/example/structs"
)

func main() {

	_, err := gostruct2graphql.ParseStructList(structs.TestSlice{})
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

}
