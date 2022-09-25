package main

import (
	"biodata/controllers"
	"fmt"
	"os"
)

func main() {

	var absen string

	if len(os.Args) < 2 {
		fmt.Println("Please provide parameter")
		os.Exit(1)
	} else {
		absen = os.Args[1]
	}

	coba := controllers.GetBiodata(absen)

	fmt.Println(coba)

}
